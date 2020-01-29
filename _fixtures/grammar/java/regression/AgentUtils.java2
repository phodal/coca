/*
 *	Copyright © 2015 Zhejiang SKT Science Technology Development Co., Ltd. All rights reserved.
 *	浙江斯凯特科技发展有限公司 版权所有
 *	http://www.28844.com
 */
package com.flycms.core.utils;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Open source house, All rights reserved
 * 开发公司：28844.com<br/>
 * 版权：开源中国<br/>
 *
 * 移动设置判断
 * 
 * @author 孙开飞 201403020
 * 
 */
public class AgentUtils {
	public static boolean choose(String userAgent) {
		if (userAgent.indexOf("Noki") > -1
				|| // Nokia phones and emulators
				userAgent.indexOf("Eric") > -1
				|| // Ericsson WAP phones and emulators
				userAgent.indexOf("WapI") > -1
				|| // Ericsson WapIDE 2.0
				userAgent.indexOf("MC21") > -1
				|| // Ericsson MC218
				userAgent.indexOf("AUR") > -1
				|| // Ericsson R320
				userAgent.indexOf("R380") > -1
				|| // Ericsson R380
				userAgent.indexOf("UP.B") > -1
				|| // UP.Browser
				userAgent.indexOf("WinW") > -1
				|| // WinWAP browser
				userAgent.indexOf("UPG1") > -1
				|| // UP.SDK 4.0
				userAgent.indexOf("upsi") > -1
				|| // another kind of UP.Browser
				userAgent.indexOf("QWAP") > -1
				|| // unknown QWAPPER browser
				userAgent.indexOf("Jigs") > -1
				|| // unknown JigSaw browser
				userAgent.indexOf("Java") > -1
				|| // unknown Java based browser
				userAgent.indexOf("Alca") > -1
				|| // unknown Alcatel-BE3 browser (UP based)
				userAgent.indexOf("MITS") > -1
				|| // unknown Mitsubishi browser
				userAgent.indexOf("MOT-") > -1
				|| // unknown browser (UP based)
				userAgent.indexOf("My S") > -1
				|| // unknown Ericsson devkit browser
				userAgent.indexOf("WAPJ") > -1
				|| // Virtual WAPJAG www.wapjag.de
				userAgent.indexOf("fetc") > -1
				|| // fetchpage.cgi Perl script from www.wapcab.de
				userAgent.indexOf("ALAV") > -1
				|| // yet another unknown UP based browser
				userAgent.indexOf("Wapa") > -1
				|| // another unknown browser (Web based "Wapalyzer")
				userAgent.indexOf("UCWEB") > -1
				|| // another unknown browser (Web based "Wapalyzer")
				userAgent.indexOf("BlackBerry") > -1
				|| // another unknown browser (Web based "Wapalyzer")
				userAgent.indexOf("J2ME") > -1
				|| // another unknown browser (Web based "Wapalyzer")
				userAgent.indexOf("Oper") > -1
				|| userAgent.indexOf("Phone") > -1
				|| userAgent.indexOf("nokia") > -1
				|| userAgent.indexOf("symbian") > -1
				|| userAgent.indexOf("iPhone") > -1 
				|| // 苹果手机ios端
				userAgent.indexOf("iPad") > -1 
				|| // iPad端
				userAgent.indexOf("iPod") > -1 
				|| // iPod端
				userAgent.indexOf("Android") > -1)

		{
			return true;
		} else {
			return false;
		}
	}

	/**
	 * 获取客户端操作系统信息，目前只匹配Win 7、WinXP、Win2003、Win2000、MAC、WinNT、Linux、Mac68k、Win9x
	 * 
	 * @param userAgent
	 *            request.getHeader("user-agent")的返回值
	 * @return
	 */
	public static String getClientOS(String userAgent) {
		String cos = "unknow os";
		String regular=".*(Windows NT 6\\.1).*";
		Pattern p = Pattern.compile(regular);
		Matcher m = p.matcher(userAgent);
		if (m.find()) {
			cos = "Win 7";
			return cos;
		}
		regular=".*(Windows NT 5\\.1|Windows XP).*";
		p = Pattern.compile(regular);
		m = p.matcher(userAgent);
		if (m.find()) {
			cos = "WinXP";
			return cos;
		}
		regular=".*(Windows NT 5\\.2).*";
		p = Pattern.compile(regular);
		m = p.matcher(userAgent);
		if (m.find()) {
			cos = "Win2003";
			return cos;
		}
		regular=".*(Win2000|Windows 2000|Windows NT 5\\.0).*";
		p = Pattern.compile(regular);
		m = p.matcher(userAgent);
		if (m.find()) {
			cos = "Win2000";
			return cos;
		}
		regular=".*(Mac|apple|MacOS8).*";
		p = Pattern.compile(regular);
		m = p.matcher(userAgent);
		if (m.find()) {
			cos = "MAC";
			return cos;
		}
		regular=".*(WinNT|Windows NT).*";
		p = Pattern.compile(regular);
		m = p.matcher(userAgent);
		if (m.find()) {
			cos = "WinNT";
			return cos;
		}
		regular=".*Linux.*";
		p = Pattern.compile(regular);
		m = p.matcher(userAgent);
		if (m.find()) {
			cos = "Linux";
			return cos;
		}
		regular=".*(68k|68000).*";
		p = Pattern.compile(regular);
		m = p.matcher(userAgent);
		if (m.find()) {
			cos = "Mac68k";
			return cos;
		}
		regular=".*(9x 4.90|Win9(5|8)|Windows 9(5|8)|95/NT|Win32|32bit).*";
		p = Pattern.compile(regular);
		m = p.matcher(userAgent);
		if (m.find()) {
			cos = "Win9x";
			return cos;
		}

		return cos;
	}
}
