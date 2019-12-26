package com.ilscipio.scipio.accounting.external.datev;

import java.lang.reflect.Constructor;
import java.util.Iterator;
import java.util.Map;

import org.ofbiz.base.util.Debug;
import org.ofbiz.entity.Delegator;
import org.ofbiz.entity.GenericValue;

import com.ilscipio.scipio.accounting.external.BaseOperationResults;
import com.ilscipio.scipio.accounting.external.BaseOperationStats;
import com.ilscipio.scipio.accounting.external.BaseOperationStats.NotificationLevel;
import com.ilscipio.scipio.accounting.external.BaseOperationStats.NotificationScope;
import com.ilscipio.scipio.accounting.external.BaseOperationStats.Stat;
import com.ilscipio.scipio.accounting.external.datev.category.AbstractDatevDataCategory;

public class DatevHelper {
    private static final Debug.OfbizLogger module = Debug.getOfbizLogger(java.lang.invoke.MethodHandles.lookup().lookupClass());

    private final String orgPartyId;
    private final BaseOperationStats stats;
    private final BaseOperationResults results;
    private final AbstractDatevDataCategory dataCategoryImpl;
    private final GenericValue dataCategory;
    private final GenericValue dataCategorySettings;

    public DatevHelper(Delegator delegator, String orgPartyId, GenericValue dataCategory) throws DatevException {
        this.orgPartyId = orgPartyId;
        try {
            this.dataCategory = dataCategory;
            this.dataCategorySettings = dataCategory.getRelatedOne("DatevGeneralSetting", true);
            @SuppressWarnings("unchecked")
            Class<? extends AbstractDatevDataCategory> dataCategoryClass = (Class<? extends AbstractDatevDataCategory>) Class.forName(dataCategory.getString("dataCategoryClass"));
            Constructor<? extends AbstractDatevDataCategory> datevDataCategoryConstructor = dataCategoryClass.getConstructor(Delegator.class, DatevHelper.class);
            this.dataCategoryImpl = datevDataCategoryConstructor.newInstance(delegator, this);
            if (dataCategoryImpl.getOperationStatsClass() != null)
                this.stats = dataCategoryImpl.getOperationStatsClass().newInstance();
            else
                this.stats = BaseOperationStats.class.newInstance();
            if (dataCategoryImpl.getOperationResultsClass() != null)
                this.results = dataCategoryImpl.getOperationResultsClass().newInstance();
            else
                this.results = BaseOperationResults.class.newInstance();
            if (Debug.isOn(Debug.VERBOSE)) {
                Debug.logInfo("Datev helper succesfully initialized.", module);
            }
        } catch (Exception e) {
            throw new DatevException("Internal error. Cannot initialize DATEV helper.");
        }

    }

    public boolean hasFatalNotification() {
        if (stats != null) {
            for (Stat stat : stats.getStats()) {
                if (stat.getLevel().equals(NotificationLevel.FATAL))
                    return true;
            }
        }
        return false;
    }

    public void processRecord(int index, Map<String, String> recordMap) throws DatevException {
        dataCategoryImpl.processRecord(index, recordMap);
    }

    public String[] getFieldNames() throws DatevException {
        return dataCategoryImpl.getDatevFieldNames();
    }

    public boolean validateField(Object o, String value) throws DatevException {
        if (o instanceof String)
            return dataCategoryImpl.validateField(String.valueOf(o), value);
        else if (o instanceof Integer)
            return dataCategoryImpl.validateField((Integer) o, value);
        return false;
    }

    public String getOrgPartyId() {
        return orgPartyId;
    }

    public boolean isMetaHeader(Iterator<String> metaHeader) throws DatevException {
        return dataCategoryImpl.isMetaHeader(metaHeader);
    }

    public BaseOperationResults getResults() {
        return results;
    }

    public BaseOperationStats getStats() {
        return stats;
    }

    public GenericValue getDataCategory() {
        return dataCategory;
    }

    public GenericValue getDataCategorySettings() {
        return dataCategorySettings;
    }

    public void addRecordStat(String message, NotificationLevel level, int position, Map<String, String> value, boolean valid) {
        addStat(stats.new RecordStat(message, level, position, value, valid));
    }

    public void addStat(String message, NotificationScope scope, NotificationLevel level) {
        addStat(stats.new Stat(message, scope, level));
    }

    private void addStat(Stat stat) {
        stats.addStat(stat);
    }

}
