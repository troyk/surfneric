package surfneric

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {

}

func TestPlacer(t *testing.T) {

	Convey("request", t, func() {
		rjasonsmith2010 := &Request{CaseYear: "2010", LastName: "Smith", FirstName: "Jason"}
		//rjasonsmith2010URL := strings.Replace(rjasonsmith2010.ToSearchURL(), siteURL, "", 1)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("TS-REQ", r.URL.Path)
			switch r.URL.Path {
			case "/CaseCalSearch/CaseIndex2/CaseIndex_list.asp":
				io.WriteString(w, jasonsmith2010)
			case "/foo":
				http.Error(w, "Unimplemented", 500)
				return
			}
		}))
		siteURL = ts.URL

		defer ts.Close()
		// test texst server harness
		//So(rjasonsmith2010.ToSearchURL(), ShouldEndWith, rjasonsmith2010URL)
		//fmt.Println(rjasonsmith2010URL)
		// fmt.Println(ts.URL)
		// fmt.Println()

		Convey("search found", func() {
			cases, err := Search(rjasonsmith2010)
			//fmt.Println("BODY", rjasonsmith2010.body)
			So(err, ShouldBeNil)
			So(cases, ShouldNotBeNil)
			So(len(cases), ShouldEqual, 1)
			So(cases[0].Number, ShouldEqual, "62-102830")
		})

		Convey("request url params", func() {
			r := &Request{CaseYear: "2010", LastName: " Stark", FirstName: " Tony"}
			url := r.ToSearchURL()
			So(url, ShouldContainSubstring, "s_F4="+sanitizeParam(r.LastName))
			So(url, ShouldContainSubstring, "s_F5="+sanitizeParam(r.FirstName))
			So(url, ShouldContainSubstring, "s_F7="+sanitizeParam(r.CaseYear))
		})

		// r := &Request{CaseYear: "2015", FirstName: "Homer", LastName:"Simpson"}
		//
		// cases := r.parseBody()
		//
		// r.CaseYear = "2014"
		// cases = r.parseBody()

		//So(Templates.Lookup("update_user_template.sql"), ShouldNotBeNil)
		//So(Templates.Lookup("update_user_named.sql"), ShouldNotBeNil)

		// Convey("arg", func() {
		// 	s := SQL("update_user_template.sql").Arg("id", "55").Arg("name", "elmer")
		// 	_, args := s.ToSQL()
		// 	So(args[0], ShouldEqual, "elmer")
		// 	So(args[1], ShouldEqual, "55")
		// })
	})
}

const jasonsmith2010 = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html><!-- InstanceBegin template="/Templates/html-template-calendars.dwt" codeOutsideHTMLIsLocked="false" -->
<head>
<meta http-equiv="X-UA-Compatible" content="IE=EmulateIE7" >
<meta http-equiv="content-type" content="text/html; charset=windows-1252">
<title>Superior Court of California, County of Placer</title>
<link rel="stylesheet" type="text/css" href="../../css/print.css" media="print" />

<link href="../../css/html5-template.css" rel="stylesheet" type="text/css"><!--[if lt IE 9]>
<script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
<![endif]-->
<link href="../../css/hefs.css" rel="stylesheet" type="text/css"><!--email contact pages-->
<link href="../../css/Spry-UI-1.7/css/Menu/basic/SpryMenuBasic.css" rel="stylesheet" type="text/css">
<script src="../../css/Spry-UI-1.7/includes/SpryDOMUtils.js" type="text/javascript"></script>
<script src="../../css/Spry-UI-1.7/includes/SpryDOMEffects.js" type="text/javascript"></script>
<script src="../../css/Spry-UI-1.7/includes/SpryWidget.js" type="text/javascript"></script>
<script src="../../css/Spry-UI-1.7/includes/SpryMenu.js" type="text/javascript"></script>
<script src="../../css/Spry-UI-1.7/includes/plugins/MenuBar2/SpryMenuBarKeyNavigationPlugin.js" type="text/javascript"></script>
<script src="../../css/Spry-UI-1.7/includes/plugins/MenuBar2/SpryMenuBarIEWorkaroundsPlugin.js" type="text/javascript"></script>
<style type="text/css">
<!--
.style1 {color: #FFFFFF}
.style3 {font-size: small}
-->
/* BeginOAWidget_Instance_2141544: #MenuBar */
/* Settable values for skinning a Basic menu via presets. If presets are not sufficient, most skinning should be done in
	these rules, with the exception of the images used for down or right pointing arrows, which are in the file SpryMenuBasic.css

	 These assume the following widget classes for menu layout (set in a preset)
	.MenuBar - Applies to all menubars - default is horizontal bar, all submenus are vertical - 2nd level subs and beyond are pull-right.
	.MenuBarVertical - vertical main bar; all submenus are pull-right.

	You can also pass in extra classnames to set your desired top level menu bar layout. Normally, these are set by using a preset.
	They only apply to horizontal menu bars:
		MenuBarLeftShrink - The menu bar will be horizontally 'shrinkwrapped' to be just large enough to hold its items, and left aligned
		MenuBarRightShrink - Just like MenuBarLeftShrink, but right aligned
		MenuBarFixedLeft - Fixed at a specified width set in the rule '.MenuBarFixedLeft', and left aligned.
		MenuBarFixedCentered -  - Fixed at a specified width set in the rule '.MenuBarFixedCentered',
						and centered in its parent container.
		MenuBarFullwidth - Grows to fill its parent container width.

	In general, all rules specified in this file are prefixed by #MenuBar so they only apply to instances of the widget inserted along
	with the rules. This permits use of multiple MenuBarBasic widgets on the same page with different layouts. Because of IE6 limitations,
	there are a few rules where this was not possible. Those rules are so noted in comments.

*/

#MenuBar  {
	background-color:#0075be;
	font-family: Verdana;/*, Geneva, sans-serif;  Specify fonts on on MenuBar and subMenu MenuItemContainer, so MenuItemContainer,
												MenuItem, and MenuItemLabel
												at a given level all use same definition for ems.
												Note that this means the size is also inherited to child submenus,
												so use caution in using relative sizes other than
												100% on submenu fonts. */
	font-weight: bolder;
	font-size: 10px;
	font-style: normal;
	padding:0;
	border-color: #ffffff #ffffff #003366 #ffffff;
	border-width:0px;
	border-style: none none solid none;
}
/* Caution: because ID+class selectors do not work properly in IE6, but we want to restrict these rules to just this
widget instance, we have used string-concatenated classnames for our selectors for the layout type of the menubar
in this section. These have very low specificity, so be careful not to accidentally override them. */

.MenuBar br { /* using just a class so it has same specificity as the ".MenuBarFixedCentered br" rule bleow */
	display:none;
}
.MenuBarLeftShrink {
	float: left; /* shrink to content, as well as float the MenuBar */
	width: auto;
}
.MenuBarRightShrink {
	float: right; /* shrink to content, as well as float the MenuBar */
	width: auto;
}
.MenuBarFixedLeft {
	float: left;
	width: 800px;
}
.MenuBarFixedCentered {
	float: none;
	width: 800px;
	margin-left:auto;
	margin-right:auto;
}
.MenuBarFixedCentered br {
	clear:both;
	display:block;
}
.MenuBarFixedCentered .SubMenu br {
	display:none;
}
.MenuBarFullwidth {
	float: left;
	width: 100%;
}

/* Top level menubar items - these actually apply to all items, and get overridden for 1st or successive level submenus */
#MenuBar  .MenuItemContainer {
	padding: 0px 0px 0px 0px;
	margin: 0; 	/* Zero out margin  on the item containers. The MenuItem is the active hover area.
				For most items, we have to do top or bottom padding or borders only on the MenuItem
				or a child so we keep the entire submenu tiled with items.
				Setting this to 0 avoids "dead spots" for hovering. */
}
#MenuBar  .MenuItem {
	padding: 0px 5px 0px 0px;
	background-color:#0075be;
	border-width:0px;
	border-color: #cccccc #ffffff #000000 #ffffff;
	border-style: none solid solid solid;
}
#MenuBar  .MenuItemFirst {
	border-style: none none none none;
}
#MenuBar .MenuItemLast {
	border-style: none solid none none;
}

#MenuBar  .MenuItem  .MenuItemLabel{
	text-align:center;
	line-height:1.4em;
	color:#ffffff;
	background-color:#0075be;
	padding: 6px 15px 6px 10px;
	width: 10em;
	width:auto;
}
.SpryIsIE6 #MenuBar  .MenuItem  .MenuItemLabel{
	width:1em; /* Equivalent to min-width in modern browsers */
}

/* First level submenu items */
#MenuBar .SubMenu  .MenuItem {
	font-family: Arial, Helvetica, sans-serif;
	font-weight: bold;
	font-size: 10px;
	font-style: normal;
	background-color:#0075be;
	padding:0px 2px 0px 0px;
	border-width:0px;
	border-color: #000000 #000000 #000000 #000000;
	/* Border styles are overriden by first and last items */
	border-style: solid solid none solid;
}
#MenuBar  .SubMenu .MenuItemFirst {
	border-style: solid solid none solid;
}
#MenuBar  .SubMenu .MenuItemFirst .MenuItemLabel{
	padding-top: 6px;
}
#MenuBar .SubMenu .MenuItemLast {
	border-style: solid solid solid solid;
}
#MenuBar .SubMenu .MenuItemLast .MenuItemLabel{
	padding-bottom: 6px;
}
#MenuBar .SubMenu .MenuItem .MenuItemLabel{
	text-align:left;
	line-height:1em;
	background-color:#0075be;
	color:#ffffff;
	width: 7em;
	width:auto;
	padding-top: 6px;
	padding-right: 24px;/*Increased padding from 12px to 24px.  added more space before submenu icon */
	padding-bottom: 6px;
	padding-left: 5px;
}

/* Hover states for containers, items and labels */
#MenuBar .MenuItemHover {
	background-color: #008fe6;
	border-color: #cccccc #cccccc #cccccc #cccccc;
}

#MenuBar .MenuItemWithSubMenu.MenuItemHover .MenuItemLabel{
	background-color: #008fe6; /* consider exposing this prop separately*/
	color: #ffffff;
}
#MenuBar .MenuItemHover .MenuItemLabel{
	background-color: #008fe6;
	color: #ffffff;
}
#MenuBar .SubMenu .MenuItemHover {
	background-color: #008fe6;
	border-color: #000000 #000000 #000000 #000000;
}

#MenuBar .SubMenu .MenuItemHover .MenuItemLabel{
	background-color: #008fe6;
	color: #ffffff;
}
/* Submenu properties -- First level of submenus */
#MenuBar .SubMenuVisible {
	background-color: #0075be;
	min-width:100%;  /* This keeps the menu from being skinnier than the parent MenuItemContainer - nice to have but not available on ie6 */
	border-color: #ffffff #ffffff #ffffff #ffffff;
	border-width:0px;
	border-style: none none none none;
}
#MenuBar.MenuBar .SubMenuVisible {/* For Horizontal menubar only */
	top: 100%;	/* 100% is at the bottom of parent menuItemContainer */
	left:0px; /* 'left' may need tuning depending upon borders or padding applied to menubar MenuItemContainer or MenuItem,
					and your personal taste.
					0px will left align the dropdown with the content area of the MenuItemContainer. Assuming you keep the margins 0
					on MenuItemContainer and MenuItem on the parent
					menubar, making this equal the sum of the MenuItemContainer & MenuItem padding-left will align
					the dropdown with the left of the menu item label.*/
	z-index:10;
}
#MenuBar.MenuBarVertical .SubMenuVisible {
	top: 0px;
	left:100%;
	min-width:0px; /* Do not neeed to match width to parent MenuItemContainer - items will prevent total collapse */
}
/* Submenu properties -- Second level submenu and beyond - these are visible descendents of .MenuLevel1 */
#MenuBar .MenuLevel1 .SubMenuVisible {
	background-color: #0075be;
	min-width:0px; /* Do not neeed to match width to parent MenuItemContainer - items will prevent total collapse*/
	top: 0px;	/* If desired, you can move this down a smidge to separate top item''s submenu from menubar -
				that is really only needed for submenu on first item of MenuLevel1, or you can make it negative to make submenu more
				vertically 'centered' on its invoking item */
	left:100%; /* If you want to shift the submenu left to partially cover its invoking item, you can add a margin-left with a
				negative value to this rule. Alternatively, if you use fixed-width items, you can change this left value
				to use px or ems to get the offset you want. */
}
/* IE6 rules - you can delete these if you do not want to support IE6 */

/* A note about multiple classes in IE6.
 * Some of the rules above use multiple class names on an element for selection, such as "hover" (MenuItemHover) and "has a subMenu" (MenuItemWithSubMenu),
 * giving the selector '.MenuItemWithSubMenu.MenuItemHover'.
 * Unfortunately IE6 does not support using mutiple classnames in a selector for an element. For a selector such as '.foo.bar.baz', IE6 ignores
 * all but the final classname (here, '.baz'), and sets the specificity accordingly, counting just one of those classs as significant. To get around this
 * problem, we use the plugin in SpryMenuBarIEWorkaroundsPlugin.js to generate compound classnames for IE6, such as 'MenuItemWithSubMenuHover'.
 * Since there are a lot of these needed, the plugin does not generate the extra classes for modern browsers, and we use the CSS2 style mutltiple class
 * syntax for that. Since IE6 both applies rules where
 * it should not, and gets the specificity wrong too, we have to order rules carefully, so the rule misapplied in IE6 can be overridden.
 * So, we put the multiple class rule first. IE6 will mistakenly apply this rule.  We follow this with the single-class rule that it would
 * mistakenly override, making sure the  misinterpreted IE6 specificity is the same as the single-class selector, so the latter wins.
 * We then create a copy of the multiple class rule, adding a '.SpryIsIE6' class as context, and making sure the specificity for
 * the selector is high enough to beat the single-class rule in the "both classes match" case. We place the IE6 rule at the end of the
 * css style block to make it easy to delete if you want to drop IE6 support.
 * If you decide you do not need IE6 support, you can get rid of these, as well as the inclusion of the SpryMenuBarIEWorkaroundsPlugin.js script.
 * The 'SpryIsIE6' class is placed on the HTML element by  the script in SpryMenuBarIEWorkaroundsPlugin.js if the browser is Internet Explorer 6. This avoids the necessity of IE conditional comments for these rules.
 */
.SpryIsIE6 #MenuBar .MenuBarView .MenuItemWithSubMenuHover .MenuItemLabel /* IE6 selector  */{
	background-color: #008fe6; /* consider exposing this prop separately*/
	color: #ffffff;
}
.SpryIsIE6 #MenuBar .MenuBarView .SubMenu .MenuItemWithSubMenuHover .MenuItemLabel/* IE6 selector  */{
	background-color: #008fe6; /* consider exposing this prop separately*/
	color: #ffffff;
}
.SpryIsIE6 #MenuBar .SubMenu .SubMenu  /* IE6 selector  */{
	margin-left: -0px; /* Compensates for at least part of an IE6 "double padding" version of the "double margin" bug */
}


/* EndOAWidget_Instance_2141544 */

</style>
<script type="text/xml">
<!--
<oa:widgets>
  <oa:widget wid="2141544" binding="#MenuBar" />
</oa:widgets>
-->
</script>
<meta content="CodeCharge Studio 4.2.00.040" name="GENERATOR">
<script language="JavaScript" type="text/javascript">
//Begin CCS script
//Include Common JSFunctions @1-5C08C693
</script>
<script language="JavaScript" src="ClientI18N.asp?file=Functions.js&amp;locale=en" type="text/javascript" charset="utf-8"></script>
<script language="JavaScript" src="ClientI18N.asp?file=DatePicker.js&amp;locale=en" type="text/javascript" charset="utf-8"></script>
<script language="JavaScript" type="text/javascript">
//End Include Common JSFunctions

//Date Picker Object Definitions @1-C180666D

var AttySearch_DatePicker_s_F4 = new Object();
AttySearch_DatePicker_s_F4.format           = "ShortDate";
AttySearch_DatePicker_s_F4.style            = "Styles/COURTS/Style.css";
AttySearch_DatePicker_s_F4.relativePathPart = "";
AttySearch_DatePicker_s_F4.themeVersion = "3.0";

var AttySearch_DatePicker_s_EndDate1 = new Object();
AttySearch_DatePicker_s_EndDate1.format           = "ShortDate";
AttySearch_DatePicker_s_EndDate1.style            = "Styles/COURTS/Style.css";
AttySearch_DatePicker_s_EndDate1.relativePathPart = "";
AttySearch_DatePicker_s_EndDate1.themeVersion = "3.0";

//End Date Picker Object Definitions
//Date Picker Object Definitions @1-22707A4D

var PublicSearch_DatePicker_s_F4 = new Object();
PublicSearch_DatePicker_s_F4.format           = "ShortDate";
PublicSearch_DatePicker_s_F4.style            = "Styles/COURTS/Style.css";
PublicSearch_DatePicker_s_F4.relativePathPart = "";
PublicSearch_DatePicker_s_F4.themeVersion = "3.0";

//End Date Picker Object Definitions

//bind_events @1-BF0EB765
function bind_events() {
}
//End bind_events

window.onload = bind_events; //Assign bind_events @1-19F7B649

//End CCS script
</script>
<link href="Styles/COURTS/Style_doctype.css" type="text/css" rel="stylesheet">
</head>
<body>
<!-- Courts Logo and American Flag -->
<div class="container">
  <header>
  <div class="hleft"><a href="#"><img src="../../images/banner1.jpg" alt="Courts Banner Logo Image" width="534" height="76" id="Insert_logo" style="background: #FFFFFF; display:block;" /></a> </div>
  <div class="hright"><img src="../../images/seal2a.jpg" width="214" height="76" alt="header right"> </div>

  </header>

<!-- Top Menu System -->
<div class="topmenu">
<!--SPRY MENU-->
<ul id="MenuBar">
    <li class="MenuGrouping"> <a href="../../index.html">Home</a></li>
<!--GENERAL INFO-->
    <li class="MenuGrouping"> <a href="#">General Info</a>
      <ul><!--Start Submenu Under General Info-->
        <li> <a href="../../general_information.html">General Introduction</a></li>
        <li> <a href="../../contact/contact_us.html">Contact Us</a></li>
        <li> <a href="../../court-calendars_1.html">Court Calendar</a></li>
        <li> <a href="../../hr/hr_intro.html">Employment</a></li>
        <li> <a href="../../fee-schedules.html">Fees</a></li>
        <li> <a href="../../court-forms.html">Forms</a></li>
        <li> <a href="../../calendar-matrix.html">Calendar Matrix</a></li>
        <li> <a href="../../records-request.html">Record / Copy Request</a></li>
        <li> <a href="../../court-history.html">History of Placer Court</a></li>
        <li> <a href="../../court-hours.html">Hours and Holidays</a></li>
        <li> <a href="../../calendar-matrix.html">Judicial Directory</a></li>
        <li> <a href="../../court-lep-plan.html">LEP Plan</a></li>
        <li> <a href="../../local-rules.html">Local Rules</a></li>
        <li> <a href="../../locations.html">Locations</a></li>
        <li> <a href="../../courts-terms-of-use.html">Web site Terms of Use</a></li>
      </ul><!--Close Submenu Under General Info-->
    </li><!--CLOSE GENERAL INFO MENU GROUP-->

<!--CIVIL-->
    <li class="MenuGrouping"> <a href="#" class="MenuBarItemSubmenu">Civil</a>
      <ul><!--Start Submenu-->
        <li><a href="../../civil/civil_intro.html">Civil Introduction</a></li>
        <li><a href="../../court-calendars_1.html">Court Calendar</a></li>
        <li><a href="../../civil/civil_forms.html">Forms</a></li>
        <li><a href="../../locations.html">Locations</a></li>
      	<li><a href="../../tentative_rulings/tentative_rulings_intro.html">Tentative Rulings</a></li>
      	<li><a href="../../civil/civil_adr.html">ADR (Alternatives to Trial)</a></li>
      	<li><a href="../../fee-schedules.html">Fees</a></li>
      	<li><a href="../../civil/civil_appeals.html">Appeals</a></li>
      	<li><a href="../../civil/civil_harassment.html">Harassment</a></li>
      	<li><a href="../../civil/civil_elder-abuse.html">Elder &amp; Dependent Abuse</a></li>
      	<li><a href="../../civil/civil_emancipation.html">Emancipation</a></li>
      	<li><a href="#">Landlord / Tenant</a>
        	<ul><!--Start Submenu-->
            <li><a href="../../civil/civil_landlord-tenant.html">Landlord / Tenant Introduction</a></li>
            <li><a href="../../civil/civil_landlord_req.html">Landlord Requirements</a></li>
        	<li><a href="../../civil/civil_tentant_req.html">Tenant Requirements</a></li>
            </ul></li><!--Close Submenu and Group-->
        <li><a href="../../civil/civil_name-change.html">Name Change</a></li>
      	<li><a href="../../civil/civil_workplace-violence.html">Workplace Violence</a></li>
     	<li><a href="#">Small Claims</a>
        	<ul><!--Start Submenu-->
       	 	<li><a href="../../civil/small-claims_help.html">Who Can Help?</a></li>
        	</ul></li><!--Close Submenu and Group-->
       	<li><a href="#">Interpreter Information</a>
        	<ul><!--Start Submenu-->
            <li><a href="../../interpreters-reporters/interpreters_information.html">Interpreter Introduction</a></li>
            <li><a href="../../interpreters-reporters/interpreters_russian.html">Russian</a></li>
        	<li><a href="../../interpreters-reporters/interpreters_simplified-chinese.html">Simplified Chinese</a></li>
        	<li><a href="../../interpreters-reporters/interpreters_spanish.html">Spanish</a></li>
        	<li><a href="../../interpreters-reporters/interpreters_traditional-chinese.html">Traditional Chinese</a></li>
            </ul></li><!--Close Submenu and Group-->
        <li><a href="#">Reporter Information</a>
        	<ul><!--Start Submenu-->
            <li><a href="../../interpreters-reporters/reporters_transcript_request.html">Transcript Request</a></li>
        	<li><a href="../../interpreters-reporters/reporters_matters.html">Reporters in Civil Matters</a></li>
            </ul></li><!--Close Submenu and Group-->
         <li><a href="../../comments/civil_contact.html">Contact Us</a></li>
     	</ul><!--Close Civil Submenu-->
     </li><!--CLOSE CIVIL MENU GROUP-->

<!--CRIMINAL-->
   <li class="MenuGrouping"> <a href="#" class="MenuBarItemSubmenu">Criminal</a>
      <ul><!--Start Submenu-->
        <li><a href='../../criminal/criminal_intro.html'>Criminal Introduction</a></li>
      	<li><a href='../../court-calendars_1.html'>Court Calendar</a></li>
      	<li><a href='../../criminal/criminal_forms.html'>Forms</a></li>
      	<li><a href='../../locations.html'>Locations</a></li>
      	<li><a href='../../fee-schedules.html'>Bail Schedule</a></li>
      	<li><a href="#">Appeals</a>
          	<ul><!--Start Submenu-->
            <li><a href='../../criminal/criminal_appeals_intro.html'>Appeals Introduction</a></li>
            <li><a href='../../criminal/criminal_appeals_nature.html'>Nature of an Appeal</a></li>
            <li><a href='../../criminal/criminal_appeals_parties.html'>Parties</a></li>
            <li><a href='../../criminal/criminal_appeals_notice.html'>Notice of an Appeal</a></li>
            <li><a href='../../criminal/criminal_appeals_proposed-statememt.html'>Proposed Statement</a></li>
            <li><a href='../../criminal/criminal_appeals_settling-statement.html'>Settling the Statement</a></li>
            <li><a href='../../criminal/criminal_appeals_transfer.html'>Transfer of Appeal</a></li>
            <li><a href='../../criminal/criminal_appeals_payment.html'>Payment of your fine</a></li>
            <li><a href='../../criminal/criminal_appeals_abandonment.html'>Abandonment of Appeal</a></li>
		  	</ul></li><!--Close Submenu and Group-->
      	<li><a href='../../criminal/criminal_expungments.html'>Expungments</a></li>
      	<li><a href='../../criminal/criminal_other-contacts.html'>Other Agency Information</a></li>
      	<li><a href='../../criminal/criminal_domestic_violence.html'>Domestic Violence</a></li>
      	<li><a href='../../criminal/criminal_faq.html'>Frequently Asked Questions</a></li>
        <li><a href="#">Interpreter Information</a>
        	<ul><!--Start Submenu-->
            <li><a href="../../interpreters-reporters/interpreters_information.html">Interpreter Introduction</a></li>
            <li><a href='../../interpreters-reporters/interpreters_russian.html'>Russian</a></li>
        	<li><a href='../../interpreters-reporters/interpreters_simplified-chinese.html'>Simplified Chinese</a></li>
        	<li><a href='../../interpreters-reporters/interpreters_spanish.html'>Spanish</a></li>
        	<li><a href='../../interpreters-reporters/interpreters_traditional-chinese.html'>Traditional Chinese</a></li>
            </ul></li><!--Close Submenu and Group-->
     	<li><a href='../../interpreters-reporters/reporters_transcript_request.html'>Reporter Information</a></li>
      	<li><a href='../../comments/criminal_contact.html'>Contact Us</a></li>
	   </ul><!--Close Criminal Submenu-->
    </li><!--CLOSE CRIMINAL MENU GROUP-->

<!--FAMILY-->
  <li class="MenuGrouping"> <a href="#" class="MenuBarItemSubmenu">Family</a>
    	<ul><!--Start Submenu-->
        	<li><a href='../../family/family_intro.html'>Family Introduction</a></li>
      		<li><a href='../../court-calendars_1.html'>Court Calendar</a></li>
      		<li><a href='../../family/family_forms.html'>Forms</a></li>
      		<li><a href='../../locations.html'>Locations</a></li>
      		<li><a href='../../tentative_rulings/family-cmc.html'>Family Law CMC</a></li>
            <li><a href='#'>Legal Help Center</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../legal/legal_intro.html'>Legal Help Center Introduction</a></li>
                <li><a href='http://www.icandocs.org/'>I-CAN! Legal</a></li>
        		<li><a href='../../legal/legal_petition-custody-support.html'>Petition for Custody and Support</a></li>
        		<li><a href='../../legal/legal_paternity-action.html'>Paternity Action</a></li>
        		<li><a href='#'>Restraining Order</a>
                	<ul><!--Start Submenu-->
                    <li><a href='../../legal/legal_restraining-order.html'>Restraining Order Introduction</a></li>
                	<li><a href='../../legal/legal_restraining-order_domestic.html'>Domestic Violence</a></li>
         			<li><a href='../../legal/legal_restraining-order_harassment.html'>Harassment</a></li>
        			</ul></li><!--Close Submenu and Group-->
                <li><a href='../../legal/legal_domestic-partnership.html'>Terminate Domestic Partnership</a></li>
        		<li><a href='../../comments/legal_contact.html'>Contact Us</a></li>
        		</ul></li><!--Close Submenu and Group-->
			<li><a href='#'>Divorce</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../family/family_divorce_intro.html'>Divorce Introduction</a></li>
                <li><a href='../../family/family_divorce.html'>Divorce</a></li>
        		<li><a href='../../family/family_legal-separation.html'>Legal Separation</a></li>
        		<li><a href='../../family/family_annulment.html'>Annulment</a></li>
        		<li><a href='../../family/family_summary-dissolution.html'>Summary Dissolutions</a></li>
                </ul></li><!--Close Submenu and Group-->
            <li><a href="#">Paternity</a>
             	<ul><!--Start Submenu-->
                <li><a href='../../family/family_paternity_intro.html'>Paternity Introduction</a></li>
                <li><a href='../../family/family_paternity.html'>Paternity</a></li>
       			<li><a href='../../family/family_paternity_visitation.html'>Custody and Visitation</a></li>
        		<li><a href='../../family/family_paternity_partnerships.html'>Domestic Partnerships</a></li>
                </ul></li><!--Close Submenu and Group-->
      		<li><a href='../../family/family_restraining-order.html'>Restraining Order</a></li>
     		<li><a href='../../family/family_request-copies.html'>Request for Copies</a></li>
     		<li><a href='../../family/family_domestic-violence.html'>Domestic Violence</a></li>
      		<li><a href='#'>Child Custody Recommending Counseling</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../family/family_mediation_intro.html'>CCRC Introduction</a></li>
                <li><a href='../../family/family_mediation_mediation.html'>CCRC</a></li>
        		<li><a href='../../family/family_mediation_before-mediation.html'>Before CCRC</a></li>
        		<li><a href='../../family/family_mediation_after-mediation.html'>After CCRC</a></li>
        		<li><a href='../../family/family_mediation_evaluations.html'>Child Custody Evaluations</a></li>
        		<li><a href='../../family/family_mediation_other-information.html'>Other Important Information</a></li>
        		</ul></li><!--Close Submenu and Group-->
			<li><a href='#'>Orientation Packet</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../family/family_op_intro.html'>Orientation Packet Introduction</a></li>
                <li><a href='../../family/family_op_flow-chart.html'>CCRC Flow Chart</a></li>
       			<li><a href='../../family/family_op_intro-family-court.html'>Intro to Family Court</a></li>
        		<li><a href='../../family/family_op_how-it-works.html'>CCRC: How it works</a></li>
        		<li><a href='../../family/family_op_worksheet.html'>Parenting Plan Worksheet</a></li>
        		<li><a href='../../family/family_op_0-6.html'>Development State: 0-6 years</a></li>
        		<li><a href='../../family/family_op_6-18.html'>Development State: 6-18 years</a></li>
        		<li><a href='../../family/family_op_formal-relationship.html'>More Formal Relationship</a></li>
        		<li><a href='../../family/family_op_vm.html'>Violence and CCRC</a></li>
        		<li><a href='../../family/family_op_helping-children.html'>Helping your children</a></li>
        		<li><a href='../../family/family_op_take-notice.html'>Parents take notice</a></li>
        		<li><a href='../../family/family_op_best-interest.html'>Best interest of the child</a></li>
        		<li><a href='../../family/family_op_classes.html'>Classes and Education</a></li>
        		<li><a href='../../family/family_op_cs-3044.html'>Family Code Section 3044</a></li>
                </ul></li><!--Close Submenu and Group-->
			<li><a href='../../family/family_child-custody-eval.html'>Child Custody Evaluation</a></li>
            <li><a href='../../family/family_child-custody-eval_spanish.html'>Child Custody Evaluation - Spanish</a></li>
            <li><a href='../../family/family_continue-drop.html'>Change Court Date Procedure</a></li>
      		<li><a href='../../family/family_faq.html'>Frequently Asked Questions</a></li>
            <li><a href='#'>Reporter Information</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../interpreters-reporters/reporters_transcript_request.html'>Transcript Request</a></li>
        		<li><a href='../../interpreters-reporters/reporters_matters.html'>Reporters in Family Matters</a></li>
    			</ul></li><!--Close Submenu and Group-->
      		<li><a href='#'>Interpreter Information</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../interpreters-reporters/interpreters_information.html'>Interpreter Information</a></li>
                <li><a href='../../interpreters-reporters/interpreters_russian.html'>Russian</a></li>
        		<li><a href='../../interpreters-reporters/interpreters_simplified-chinese.html'>Simplified Chinese</a></li>
       		 	<li><a href='../../interpreters-reporters/interpreters_spanish.html'>Spanish</a></li>
       		 	<li><a href='../../interpreters-reporters/interpreters_traditional-chinese.html'>Traditional Chinese</a></li>
                </ul></li><!--Close Submenu and Group-->
			<li><a href='../../comments/family_contact.html'>Contact Us</a></li>
       </ul><!--Close Family Submenu-->
	</li><!--CLOSE FAMILY MENU GROUP-->

<!--PROBATE-->
	<li class="MenuGrouping"> <a href="#" class="MenuBarItemSubmenu">Probate</a>
    	<ul><!--Start Submenu-->
        	<li><a href='../../civil/civil_probate.html'>Probate Introduction</a></li>
     		<li><a href='../../court-calendars_1.html'>Court Calendar</a></li>
      		<li><a href='../../civil/civil_probate_forms.html'>Forms</a></li>
      		<li><a href='../../locations.html'>Locations</a></li>
      		<li><a href='../../civil/civil_probate_guardianship.html'>Guardianship</a></li>
      		<li><a href='../../civil/civil_probate_conservatorship.html'>Conservatorship</a></li>
            <li><a href='#'>Reporter Information</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../interpreters-reporters/reporters_transcript_request.html'>Transcript Request</a></li>
        		<li><a href='../../interpreters-reporters/reporters_matters.html'>Reporters in Family Matters</a></li>
    			</ul></li><!--Close Submenu and Group-->
            <li><a href='#'>Interpreter Information</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../interpreters-reporters/interpreters_information.html'>Interpreter Information</a></li>
                <li><a href='../../interpreters-reporters/interpreters_russian.html'>Russian</a></li>
        		<li><a href='../../interpreters-reporters/interpreters_simplified-chinese.html'>Simplified Chinese</a></li>
        		<li><a href='../../interpreters-reporters/interpreters_spanish.html'>Spanish</a></li>
        		<li><a href='../../interpreters-reporters/interpreters_traditional-chinese.html'>Traditional Chinese</a></li>
                </ul></li><!--Close Submenu and Group-->
              <li><a href='../../comments/civil_contact.html'>Contact Us</a></li>
         </ul><!--Close Probate Submenu-->
      </li><!--CLOSE PROBATE MENU GROUP-->

<!--JUVENILE-->
	<li class="MenuGrouping"> <a href="#">Juvenile</a>
    	<ul><!--Start Submenu-->
        	<li><a href='../../juvenile/juvenile_intro.html'>Juvenile Introduction</a></li>
			<li><a href='../../juvenile/juvenile_forms.html'>Forms</a></li>
      		<li><a href='../../locations.html'>Locations</a></li>
      		<li><a href='../../juvenile/juvenile_delinquency.html'>Juvenile Delinquency</a></li>
      		<li><a href='../../juvenile/juvenile_dependency.html'>Juvenile Dependency</a></li>
      		<li><a href='../../JJDPC/JJDPC.html'>Juvenile Justice and Delinquency Prevention Commission</a></li>
      		<li><a href='../../juvenile/juvenile_faq.html'>Frequently Asked Questions</a></li>
      		<li><a href='../../comments/juvenile_contact.html'>Contact Us</a></li>
    	</ul><!--Close Juvenile Submenu-->
	</li><!--CLOSE JUVENILE MENU GROUP-->

<!--TRAFFIC-->
	<li class="MenuGrouping"> <a href="#">Traffic</a>
    	<ul><!--Start Submenu-->
        	<li><a href='../../traffic/traffic_intro.html'>Traffic Introduction</a></li>
        	<li><a href='../../court-calendars_1.html'>Courts Calendar</a></li>
      		<li><a href='../../court-calendars_3.html'>Citation Payment Information</a></li>
      		<li><a href='../../traffic/traffic_forms.html'>Forms</a></li>
      		<li><a href='../../locations.html'>Locations</a></li>
      		<li><a href='#'>Plea</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../traffic/traffic_plea_intro.html'>Plea Introduction</a></li>
                <li><a href='../../traffic/traffic_plea_guilty.html'>Guilty - No Contest</a></li>
        		<li><a href='../../traffic/traffic_plea_not-guilty.html'>Not Guilty</a></li>
        		<li><a href='../../traffic/traffic_plea_officer-court-trial.html'>Officer Court Trial</a></li>
        		<li><a href='../../traffic/traffic_plea_trial-by-declaration.html'>Trial By Declaration</a></li>
                </ul></li><!--Close Submenu and Group-->
			<li><a href='../../traffic/traffic_traffic-school.html'>Traffic School</a></li>
      		<li><a href='../../traffic/traffic_online-payments.html'>Pay by Credit Card</a></li>
      		<li><a href='../../traffic/traffic_payment-plan.html'>Payment Plan</a></li>
      		<li><a href='../../traffic/traffic_other-agency.html'>Other Agency Information</a></li>
      		<li><a href='../../traffic/traffic_faq.html'>Frequently Asked Questions</a></li>
     		<li><a href='../../comments/traffic_contact.html'>Contact Us</a></li>
        </ul><!--Close trafic Submenu-->
	</li><!--CLOSE TRAFFIC MENU GROUP-->

<!--JURY-->
	<li class="MenuGrouping"> <a href="#">Jury</a>
    	<ul><!--Start Submenu-->
        	<li><a href='../../jury/jury_intro.html'>Jury Introduction</a></li>
   		    <li><a href='../../jury/jury_jury-online.html'>Online Jury Services</a></li>
      		<li><a href='../../locations.html'>Locations</a></li>
      		<li><a href='#'>Jury Duty Information</a>
      			<ul><!--Start Submenu-->
                	<li><a href='../../jury/jury_intro.html'>Jury Duty Introduction</a></li>
        			<li><a href='../../jury/jury_general_american-responsibility.html'>American Responsibility</a></li>
        			<li><a href='../../jury/jury_general_parking.html'>Juror Parking</a></li>
        			<li><a href='../../jury/jury_general_potential-jurors.html'>Potential Jurors</a></li>
        			<li><a href='../../jury/jury_general_postponement.html'>Postponement</a></li>
        			<li><a href='../../jury/jury_general_exemptions.html'>Exemptions</a></li>
        			<li><a href='../../jury/jury_general_job.html'>Regular Job</a></li>
        			<li><a href='../../jury/jury_general_hours.html'>Hours of Service</a></li>
        			<li><a href='../../jury/jury_general_security.html'>Court Security</a></li>
        			<li><a href='../../jury/jury_general_attire.html'>Recommended Attire</a></li>
       			 	<li><a href='../../jury/jury_general_compensation.html'>Compensation</a></li>
       				<li><a href='../../jury/jury_general_jury-trials.html'>Jury Trials</a></li>
        			<li><a href='../../jury/jury_general_process.html'>The Process</a></li>
        			<li><a href='../../jury/jury_general_trial.html'>The Trial</a></li>
        			<li><a href='../../jury/jury_general_verdict.html'>The Verdict</a></li>
        			<li><a href='../../jury/jury_general_misc.html'>Misc Information</a></li>
        			<li><a href='../../jury/jury_faq.html'>Frequently Asked Questions</a></li>
                    </ul></li><!--Close Submenu and Group-->
      		<li><a href='#'>Grand Jury</a>
				<ul><!--Start Submenu-->
                	<li><a href='../../grandjury/grandjury-fp.html'>Grand Jury Introduction</a></li>
        			<li><a href='../../grandjury/grandjury_intro.html'>What is the Grand Jury?</a></li>
        			<li><a href='../../grandjury/grandjury_forms.html'>Forms</a></li>
				<li><a href='#'>Grand Jury Reports</a>
                	<ul><!--Start Submenu-->
                    	<li><a href='../../grandjury/grandjury_reports.html'>Reports Listing</a>
          				<li><a href='../../grandjury/2011-2012/2011-2012.html'>2011 - 2012</a></li>
          				<li><a href='../../grandjury/2010-2011/2010-2011.html'>2010 - 2011</a></li>
          				<li><a href='../../grandjury/2009-2010/2009-2010.html'>2009 - 2010</a></li>
         				<li><a href='../../grandjury/2008-2009/2008-2009.html'>2008 - 2009</a></li>
         				<li><a href='../../grandjury/2007-2008/2007-2008.html'>2007 - 2008</a></li>
         				<li><a href='../../grandjury/2006-2007/2006-2007.html'>2006 - 2007</a></li>
          				<li><a href='../../grandjury/2005-2006/2005-2006.html'>2005 - 2006</a></li>
          				<li><a href='../../grandjury/2004-2005/2004-2005.html'>2004 - 2005</a></li>
          				<li><a href='../../grandjury/2003-2004/2003-2004.html'>2003 - 2004</a></li>
          				<li><a href='../../grandjury/2002-2003/2002-2003.html'>2002 - 2003</a></li>
          				<li><a href='../../grandjury/2001-2002/2001-2002.html'>2001 - 2002</a></li>
          				<li><a href='../../grandjury/2000-2001/2000-2001.html'>2000 - 2001</a></li>
          				<li><a href='../../grandjury/1999-2000/1999-2000.html'>1999 - 2000</a></li>
                        </ul></li><!--Close Submenu and Group-->
        		<li><a href='../../grandjury/grandjury_resolutions.html'>Resolutions</a></li>
        		<li><a href='../../grandjury/grandjury_faq.html'>Frequently Asked Questions</a></li>
              </ul></li><!--Close Submenu and Group-->
      		<li><a href='../../comments/jury_contact.html'>Contact Us</a></li>
        </ul><!--Close Jury Submenu-->
	</li><!--CLOSE JURY MENU GROUP-->

<!--APPEALS-->
	<li class="MenuGrouping"> <a href="#">Appeals</a>
		<ul><!--Start Submenu-->
     		<li><a href='../../criminal/criminal_appeals_intro.html'>Appeals Introduction</a></li>
      		<li><a href='../../fee-schedules.html'>Fees</a></li>
      		<li><a href='../../court-forms.html'>Forms</a></li>
      		<li><a href='../../locations.html'>Locations</a></li>
      		<li><a href='#'>Criminal Appeal</a>
				<ul><!--Start Submenu-->
                	<li><a href='../../criminal/criminal_appeals_intro.html'>Appeals Introduction</a></li>
        			<li><a href='../../criminal/criminal_appeals_nature.html'>Nature of an Appeal</a></li>
       				<li><a href='../../criminal/criminal_appeals_parties.html'>Parties</a></li>
        			<li><a href='../../criminal/criminal_appeals_notice.html'>Notice of Appeal</a></li>
        			<li><a href='../../criminal/criminal_appeals_proposed-statememt.html'>Proposed Statement</a></li>
        			<li><a href='../../criminal/criminal_appeals_settling-statement.html'>Settling the Statement</a></li>
        			<li><a href='../../criminal/criminal_appeals_transfer.html'>Transfer of Appeal</a></li>
        			<li><a href='../../criminal/criminal_appeals_payment.html'>Payment of your fine</a></li>
        			<li><a href='../../criminal/criminal_appeals_abandonment.html'>Abandonment of Appeal</a></li>
                    </ul></li><!--Close Submenu and Group-->
      		<li><a href='#'>Reporter Information</a>
            	<ul><!--Start Submenu-->
                <li><a href='../../interpreters-reporters/reporters_transcript_request.html'>Transcript Request</a></li>
        		<li><a href='../../interpreters-reporters/reporters_matters.html'>Reporters in Family Matters</a></li>
    			</ul></li><!--Close Submenu and Group-->
         </ul><!--Close Appeals Submenu-->
	</li><!--CLOSE APPEALS MENU GROUP-->
</ul><!--CLOSE WHOLE MENU SYSTEM-->

<script type="text/javascript">
// BeginOAWidget_Instance_2141544: #MenuBar
var MenuBar = new Spry.Widget.MenuBar2("#MenuBar", {
      widgetID: "MenuBar",
	  widgetClass: "MenuBar  MenuBarLeftShrink",
	  insertMenuBarBreak: true
});
// EndOAWidget_Instance_2141544
</script>
<!--*****************END SPRY MENU SYSTEM*****************-->

<!--*****************START SEARCH ENGINE*****************-->
<script type="text/javascript">

// Google Internal Site Search script- By JavaScriptKit.com (http://www.javascriptkit.com)
// For this and over 400+ free scripts, visit JavaScript Kit- http://www.javascriptkit.com/
// This notice must stay intact for use

var domainroot="www.placer.courts.ca.gov"

function Gsitesearch(curobj){
curobj.q.value="site:"+domainroot+" "+curobj.qfront.value
}

</script>

<div class="search">
<form style="margin-top:0px; margin-bottom:0px;" action="http://www.google.com/search" method="get" onsubmit="Gsitesearch(this)">
<input name="q" type="hidden" />
<input name="qfront" type="text" style="width: 130px" id="qfront" /><input type="submit" value="Search" />
<br />
<label for="qfront" style="font-size:9px"></label>
</form>
</div><!--*****************End Search Engine*****************-->

</div>
<!--- End Search Engine -->
<!-- InstanceBeginEditable name="codecharger" -->




<div id="content">
  <!--- Start of Content Area -->
  <table id="menucontent" width="100%" border="0" >
  <tr>
    <td>

<form id="CaseIndexSearch" name="CaseIndexSearch" action="CaseIndex_list.asp?s_F5=jason&amp;s_CN=&amp;s_F3=Criminal&amp;s_F7=2010&amp;s_F4=smith&amp;ccsForm=CaseIndexSearch" method="post">
  <table cellspacing="0" cellpadding="0" border="0">
    <tr>
      <td valign="top">
        <table class="Header" cellspacing="0" cellpadding="0" border="0">
          <tr>
            <td class="HeaderLeft"><img alt="" src="Styles/COURTS/Images/Spacer.gif"
              border="0"></td>
            <td class="th"><strong>Public Case Index </strong></td>
            <td class="HeaderRight"><img alt="" src="Styles/COURTS/Images/Spacer.gif"
              border="0"></td>
          </tr>
        </table>

        <table class="Record" cellspacing="0" cellpadding="0">

          <tr class="Controls">
            <td class="th">Case No.:</td>
            <td>
              <input id="CaseIndexSearchs_CN" value="" name="s_CN">
            </td>
            <td rowspan="9">
              <ul>
                <li><font
face="@Arial Unicode MS"><font size="2">Maximum number of records
                returned&nbsp;= 1,000</font></font>
                <li><font face="@Arial Unicode MS"
                size="2">Older cases use 1999 as the case year</font>
                <li><font face="@Arial Unicode MS" size="2">For best results, select the
                case type&nbsp;and enter&nbsp;the first&nbsp;and last name
</font>
                <li><font face="@Arial Unicode MS" size="2">Fields with "*" are required
                </font></li>
              </ul>
              <p><br>
              </p>
              <p>&nbsp; </p></td>
          </tr>

          <tr class="Controls">
            <td class="th" colspan="2">&nbsp;-- OR --</td>
          </tr>

          <tr class="Controls">
            <td class="th"><label for="CaseIndexSearchs_F3">Case Type:</label>
</td>
            <td>
              <select id="CaseIndexSearchs_F3" style="WIDTH: 213px" name="s_F3">
                <option value="" selected>Select Value</option>
                <OPTION VALUE="Civil">Civil</OPTION>
<OPTION VALUE="Criminal" SELECTED>Criminal</OPTION>
<OPTION VALUE="Family Law">Family Law</OPTION>
<OPTION VALUE="Probate">Probate</OPTION>
<OPTION VALUE="Small Claims">Small Claims</OPTION>
<OPTION VALUE="%">All</OPTION>

              </select>
 &nbsp;*</td>
          </tr>

          <tr class="Controls">
            <td class="th" colspan="2">&nbsp; --&nbsp;AND --</td>
          </tr>

          <tr class="Controls">
            <td class="th">Case Year:</td>
            <td>
              <input id="CaseIndexSearchs_F7" value="2010" name="s_F7">
              &nbsp;*</td>
          </tr>

          <tr class="Controls">
            <td class="th" colspan="2">&nbsp;&nbsp;-- AND --</td>
          </tr>

          <tr class="Controls">
            <td class="th"><label for="CaseIndexSearchs_F4">Last Name:</label>
</td>
            <td>
              <input id="CaseIndexSearchs_F4" maxlength="250" size="50" value="smith"
              name="s_F4">&nbsp;*
            </td>
          </tr>

          <tr class="Controls">
            <td class="th" colspan="2">&nbsp; --&nbsp;AND / OR --</td>
          </tr>

          <tr class="Controls">
            <td class="th"><label for="CaseIndexSearchs_F5">First Name:</label>
              </td>
            <td>
              <input id="CaseIndexSearchs_F5" maxlength="250" size="50" value="jason"
              name="s_F5">
            </td>
          </tr>

          <tr class="Bottom">
            <td align="right" colspan="3">

              <input class="Button" id="CaseIndexSearchButton_DoSearch" type="submit"
              alt="Search" value="Search" name="Button_DoSearch">


              <input class="Button" id="CaseIndexSearchButton1" type="submit" alt="Button1"
              value="Clear" name="Button1">
              </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</form>
<br>

<table cellspacing="0" cellpadding="0" border="0">
  <tr>
    <td valign="top">
      <table class="Header" cellspacing="0" cellpadding="0" border="0">
        <tr>
          <td class="HeaderLeft"><img alt="" src="Styles/COURTS/Images/Spacer.gif"
            border="0"></td>
          <td class="th"><strong>Search Results</strong></td>
          <td class="HeaderRight"><img alt="" src="Styles/COURTS/Images/Spacer.gif"
            border="0"></td>
        </tr>
      </table>

      <table class="Grid" cellspacing="0" cellpadding="0">
        <tr class="Caption">
          <th scope="col">
          <a href="CaseIndex_list.asp?s_F5=jason&amp;s_CN=&amp;s_F3=Criminal&amp;s_F7=2010&amp;s_F4=smith&amp;CaseIndexOrder=Sorter_F1&amp;CaseIndexDir=ASC"
id="CaseIndexSorter_F1">Case No.</a>

          </th>

          <th scope="col">
          <a href="CaseIndex_list.asp?s_F5=jason&amp;s_CN=&amp;s_F3=Criminal&amp;s_F7=2010&amp;s_F4=smith&amp;CaseIndexOrder=Sorter_F3&amp;CaseIndexDir=ASC"
id="CaseIndexSorter_F3">Case Type</a>

          </th>

          <th scope="col">
          <a href="CaseIndex_list.asp?s_F5=jason&amp;s_CN=&amp;s_F3=Criminal&amp;s_F7=2010&amp;s_F4=smith&amp;CaseIndexOrder=Sorter_F4&amp;CaseIndexDir=ASC"
id="CaseIndexSorter_F4">Last Name</a>

          </th>

          <th scope="col">
          <a href="CaseIndex_list.asp?s_F5=jason&amp;s_CN=&amp;s_F3=Criminal&amp;s_F7=2010&amp;s_F4=smith&amp;CaseIndexOrder=Sorter_F5&amp;CaseIndexDir=ASC"
id="CaseIndexSorter_F5">First Name</a>

          </th>

          <th scope="col">
          <a href="CaseIndex_list.asp?s_F5=jason&amp;s_CN=&amp;s_F3=Criminal&amp;s_F7=2010&amp;s_F4=smith&amp;CaseIndexOrder=Sorter_F8&amp;CaseIndexDir=ASC"
id="CaseIndexSorter_F8">Case Caption</a>

          &nbsp;</th>

          <th scope="col">
          <a href="CaseIndex_list.asp?s_F5=jason&amp;s_CN=&amp;s_F3=Criminal&amp;s_F7=2010&amp;s_F4=smith&amp;CaseIndexOrder=Sorter_F6&amp;CaseIndexDir=ASC"
id="CaseIndexSorter_F6">Party Type</a>

          </th>

          <th scope="col">
          <a href="CaseIndex_list.asp?s_F5=jason&amp;s_CN=&amp;s_F3=Criminal&amp;s_F7=2010&amp;s_F4=smith&amp;CaseIndexOrder=Sorter_F7&amp;CaseIndexDir=ASC"
id="CaseIndexSorter_F7">Case Date</a>

          </th>
        </tr>


        <tr class="Row">
          <td>62-102830&nbsp;</td>
          <td>Criminal&nbsp;</td>
          <td>Smith&nbsp;</td>
          <td>Jason&nbsp;</td>
          <td>&nbsp;People vs. Smith, Jason Tyler</td>
          <td>Defendant&nbsp;</td>
          <td>11/16/2010&nbsp;</td>
        </tr>


        <tr class="Footer">
          <td colspan="7">
            &nbsp;
            </td>
        </tr>
      </table>
    </td>
  </tr>
</table>



</td>
  </tr>
</table>
</div>
<!-- InstanceEndEditable -->

<!--
<div id="disclaimer">
<p class="style3">&nbsp;</p>
<p class="style3">? 2006 - 2012 Superior Court of California,
  County of Placer.&nbsp; All rights reserved<br />
  <a href="../../courts-terms-of-use.html">
    Terms Of Use</a>&nbsp; |&nbsp; Comments regarding this Web site may be directed to: <a href="../../comments/webmaster_contact.html">webmaster@placer.courts.ca.gov</a></p>
<p>&nbsp;</p>
</div>
-->

<!--***************START OF FOOTER**************-->
<footer></footer>

<footer1>
    <address>
    <p>? 2006 - 2012 Superior Court of California, County of Placer.  All rights reserved<br>
      <a href="../../courts-terms-of-use.html">Terms Of Use</a> | Site Map | Comments regarding this Web site may be directed to: webmaster@placer.courts.ca.gov</p>
    <p></p>
    </address>
  </footer1>


</body>
<!-- InstanceEnd --></html>`
