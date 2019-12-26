package org.ofbiz.base.util;

import java.io.File;
import java.net.MalformedURLException;
import java.net.URI;
import java.net.URL;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * URL Utilities - Simple Class for flexibly working with properties files
 *
 */
public final class UtilURL {

    private static final Debug.OfbizLogger module = Debug.getOfbizLogger(java.lang.invoke.MethodHandles.lookup().lookupClass());
    private static final Map<String, URL> urlMap = new ConcurrentHashMap<>();

    private UtilURL() {}

    public static <C> URL fromClass(Class<C> contextClass) {
        String resourceName = contextClass.getName();
        int dotIndex = resourceName.lastIndexOf('.');

        if (dotIndex != -1) {
            resourceName = resourceName.substring(0, dotIndex);
        }
        resourceName += ".properties";

        return fromResource(contextClass, resourceName);
    }

    /**
     * Returns a <code>URL</code> instance from a resource name. Returns
     * <code>null</code> if the resource is not found.
     * <p>This method uses various ways to locate the resource, and in all
     * cases it tests to see if the resource exists - so it
     * is very inefficient.</p>
     *
     * @param resourceName
     * @return
     */
    public static URL fromResource(String resourceName) {
        return fromResource(resourceName, null);
    }

    public static <C> URL fromResource(Class<C> contextClass, String resourceName) {
        if (contextClass == null) {
            return fromResource(resourceName, null);
        }
        return fromResource(resourceName, contextClass.getClassLoader());
    }

    /**
     * Returns a <code>URL</code> instance from a resource name. Returns
     * <code>null</code> if the resource is not found.
     * <p>This method uses various ways to locate the resource, and in all
     * cases it tests to see if the resource exists - so it
     * is very inefficient.</p>
     *
     * @param resourceName
     * @param loader
     * @return
     */
    public static URL fromResource(String resourceName, ClassLoader loader) {
        URL url = urlMap.get(resourceName);
        if (url != null) {
            try {
                return new URL(url.toString());
            } catch (MalformedURLException e) {
                Debug.logWarning(e, "Exception thrown while copying URL: ", module);
            }
        }
        if (loader == null) {
            try {
                loader = Thread.currentThread().getContextClassLoader();
            } catch (SecurityException e) {
                // Huh? The new object will be created by the current thread, so how is this any different than the previous code?
                loader = UtilURL.class.getClassLoader();
            }
        }
        url = loader.getResource(resourceName);
        if (url != null) {
            urlMap.put(resourceName, url);
            return url;
        }
        url = ClassLoader.getSystemResource(resourceName);
        if (url != null) {
            urlMap.put(resourceName, url);
            return url;
        }
        url = fromFilename(resourceName);
        if (url != null) {
            urlMap.put(resourceName, url);
            return url;
        }
        url = fromOfbizHomePath(resourceName);
        if (url != null) {
            urlMap.put(resourceName, url);
            return url;
        }
        url = fromUrlString(resourceName);
        if (url != null) {
            urlMap.put(resourceName, url);
        }
        return url;
    }

    public static URL fromFilename(String filename) {
        if (filename == null) {
            return null;
        }
        File file = new File(filename);
        URL url = null;

        try {
            if (file.exists()) {
                url = file.toURI().toURL();
            }
        } catch (java.net.MalformedURLException e) {
            Debug.logError(e, "unable to retrieve URL for file: " + filename, module);
            url = null;
        }
        return url;
    }

    public static URL fromUrlString(String urlString) {
        URL url = null;
        try {
            url = new URL(urlString);
        } catch (MalformedURLException e) {
        }

        return url;
    }

    public static URL fromOfbizHomePath(String filename) {
        String ofbizHome = System.getProperty("ofbiz.home");
        if (ofbizHome == null) {
            Debug.logWarning("No ofbiz.home property set in environment", module);
            return null;
        }
        String newFilename = ofbizHome;
        if (!newFilename.endsWith("/") && !filename.startsWith("/")) {
            newFilename = newFilename + "/";
        }
        newFilename = newFilename + filename;
        return fromFilename(newFilename);
    }

    /**
     * Gets file location (URL) relative to project root from an absolute file path.
     * NOTE: The file path is NOT a URL! This is a convenience method.
     * <p>
     * SCIPIO: NOTE: 2018-12-06: This method is modified so that it will now return null
     * if the given fileUrl is not under the project root; this is done for common sense
     * and for security concerns about the original callers of this method.
     */
    public static String getOfbizHomeRelativeLocationFromFilePath(String path) { // SCIPIO: String overload + impl moved
        if (path == null) {
            return null;
        }
        String ofbizHome = System.getProperty("ofbiz.home");
        if (path.startsWith(ofbizHome)) {
            // SCIPIO: Added length check and missing slash comparison
            if (path.length() == ofbizHome.length()) {
                return "";
            } else if (path.charAt(ofbizHome.length()) == '/') {
                // note: the +1 is to remove the leading slash
                return path.substring(ofbizHome.length()+1);
            }
        }
        // SCIPIO: Is not applicable, return null.
        //return path;
        return null;
    }

    /**
     * Gets file location (URL) relative to project root.
     * <p>
     * SCIPIO: NOTE: 2018-12-06: This method is modified so that it will now return null
     * if the given fileUrl is not under the project root; this is done for common sense
     * and for security concerns about the original callers of this method.
     */
    public static String getOfbizHomeRelativeLocation(URL fileUrl) {
        if (fileUrl == null) {
            return null;
        }
        return getOfbizHomeRelativeLocationFromFilePath(fileUrl.getPath()); // SCIPIO: refactored
    }

    /**
     * SCIPIO: Gets file location (URI) relative to project root.
     * <p>
     * SCIPIO: NOTE: 2018-12-06: This method is modified so that it will now return null
     * if the given fileUrl is not under the project root; this is done for common sense
     * and for security concerns about the original callers of this method.
     */
    public static String getOfbizHomeRelativeLocation(URI fileUrl) { // SCIPIO: URI overload
        if (fileUrl == null) {
            return null;
        }
        return getOfbizHomeRelativeLocationFromFilePath(fileUrl.getPath());

    }
}
