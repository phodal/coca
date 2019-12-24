package com.zheng.upms.server.controller.manage;

import com.baidu.unbiz.fluentvalidator.ComplexResult;
import com.baidu.unbiz.fluentvalidator.FluentValidator;
import com.baidu.unbiz.fluentvalidator.ResultCollectors;
import com.zheng.common.base.BaseController;
import com.zheng.common.validator.LengthValidator;
import com.zheng.upms.common.constant.UpmsResult;
import com.zheng.upms.common.constant.UpmsResultConstant;
import com.zheng.upms.dao.model.UpmsSystem;
import com.zheng.upms.dao.model.UpmsSystemExample;
import com.zheng.upms.rpc.api.UpmsSystemService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.apache.commons.lang.StringUtils;
import org.apache.shiro.authz.annotation.RequiresPermissions;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * 系统controller
 * Created by shuzheng on 2016/12/18.
 */
@Controller
@Api(value = "系统管理", description = "系统管理")
@RequestMapping("/manage/system")
public class UpmsSystemController extends BaseController {

	private static final Logger LOGGER = LoggerFactory.getLogger(UpmsSystemController.class);

	@Autowired
	private UpmsSystemService upmsSystemService;

	@ApiOperation(value = "系统列表")
	@RequiresPermissions("upms:system:read")
	@RequestMapping(value = "/list", method = RequestMethod.GET)
	@ResponseBody
	public Object list(
			@RequestParam(required = false, defaultValue = "0", value = "offset") int offset,
			@RequestParam(required = false, defaultValue = "10", value = "limit") int limit,
			@RequestParam(required = false, defaultValue = "", value = "search") String search,
			@RequestParam(required = false, value = "sort") String sort,
			@RequestParam(required = false, value = "order") String order) {
		UpmsSystemExample upmsSystemExample = new UpmsSystemExample();
		if (!StringUtils.isBlank(sort)
		    && !StringUtils.isBlank(order)

		 ) {
			upmsSystemExample.setOrderByClause(sort + " " + order);
		}
		if (StringUtils.isNotBlank(search)) {
			upmsSystemExample.or()
					.andTitleLike("%" + search + "%");
		}
		List<UpmsSystem> rows = upmsSystemService.selectByExampleForOffsetPage(upmsSystemExample, offset, limit);
		long total = upmsSystemService.countByExample(upmsSystemExample);
		Map<String, Object> result = new HashMap<>();
		result.put("rows", rows);
		result.put("total", total);
		return result;
	}
}
