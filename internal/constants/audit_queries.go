package constants

import ()

type AuditItem struct {
	Issue	 	string
	IssueName 	string
	IssueDef 	string
	Query 		string
	Category 	string
	ElId 		uint16
	Checked 	bool	
	IType 		string
}

type AppAuditVwData struct {
	Base 			*BaseTemplateparams
	Data 			[]AuditItem
	TextAreaStr 	string
	OldView 		bool
}

/*
var AUDIT_ITEMS = []AuditItem{{
		Issue: 		"Issue#1",	
		IssueName: 	"Duplicate Buttons",
		IssueDef: 	"select 1 as 'Issue#', 'Duplicate Buttons' as 'Issue', round(count(*)/2) as 'NbrOfIssues' from (select c.loginName, d.swap, a.*, b.totalCount as Duplicate from User c, UserCDI d, Button a inner join (select buttonNumber, parentUserCDIId, count(*) totalCount from Button group by Button.buttonNumber, Button.parentUserCDIId having count(*) >=2) b on b.buttonNumber = a.buttonNumber and b.parentUserCDIId = a.parentUserCDIId where c.userCDIId=a.parentUserCDIId and c.userCDIId=d.id) as x",
		Query: 		"select c.loginName, d.swap, a.*, b.totalCount as Duplicate from User c, UserCDI d, Button a inner join (select buttonNumber, parentUserCDIId, count(*) totalCount from Button group by Button.buttonNumber, Button.parentUserCDIId having count(*) >=2) b on b.buttonNumber = a.buttonNumber and b.parentUserCDIId = a.parentUserCDIId where c.userCDIId=a.parentUserCDIId and c.userCDIId=d.id", 
		Category: 	"Audit Query", 
		ElId: 		"audit1",
		Checked: 	true,
	},
	{
		Issue: 		"Issue#2",	
		IssueName: 	"Open Connexion lines on speakers with no listening permissions", 
		IssueDef: 	"select 2 as 'Issue#', 'Open Connexion lines on speakers with no listening permissions' as 'Issue', count(*) as 'NbrOfIssues' from (select a.loginName, a.id as userId, a.userCDIId, b.speakerNumber, b.label, b.activeStatus, c.id as resourceAORIdnoLstn from User a, UserSpeakerChannel b, ResourceAOR c where a.userCDIId=b.parentUserCDIId and b.resourceAORId=c.id and c.type='OpenConnexion' and a.accountType='real' and a.id not in (select listenersId from ResourceAORUserListenersMap d where c.id=d.resourceAORId)) as y",
		Query: 		"select a.loginName, a.id as userId, a.userCDIId, b.speakerNumber, b.label, b.activeStatus, c.id as resourceAORIdnoLstn from User a, UserSpeakerChannel b, ResourceAOR c where a.userCDIId=b.parentUserCDIId and b.resourceAORId=c.id and c.type='OpenConnexion' and a.accountType='real' and a.id not in (select listenersId from ResourceAORUserListenersMap d where c.id=d.resourceAORId)", 
		Category: 	"Audit Query", 
		ElId: 		"audit2",
		Checked: 	true,
	},
	{
		Issue: 		"Issue#3",	
		IssueName: 	"Open Connexion lines on buttons with no listening permissions", 
		IssueDef: 	"select 3 as 'Issue#', 'Open Connexion lines on buttons with no listening permissions' as 'Issue', count(*) as 'NbrOfIssues' from (select a.loginName, a.id as userId, a.userCDIId, b.buttonNumber, b.buttonLabel, c.id as resourceAORIdnoListn, e.appearance from User a, Button b, ResourceAOR c, ButtonResourceAppearance e where a.userCDIId=b.parentUserCDIId and e.resourceAORId=c.id and c.type='OpenConnexion' and a.accountType='real' and e.parentButtonId=b.id and a.id not in (select listenersId from ResourceAORUserListenersMap d where c.id=d.resourceAORId)) as x",
		Query: 		"select a.loginName, a.id as userId, a.userCDIId, b.buttonNumber, b.buttonLabel, c.id as resourceAORIdnoListn, e.appearance from User a, Button b, ResourceAOR c, ButtonResourceAppearance e where a.userCDIId=b.parentUserCDIId and e.resourceAORId=c.id and c.type='OpenConnexion' and a.accountType='real' and e.parentButtonId=b.id and a.id not in (select listenersId from ResourceAORUserListenersMap d where c.id=d.resourceAORId)",		
		Category: 	"Audit Query", 
		ElId: 		"audit3",
		Checked: 	true,
	},
	{
		Issue: 		"Issue#4",	
		IssueName: 	"UserCDI with no User id", 
		IssueDef: 	"select 4 as 'Issue#', 'UserCDI with no User id' as 'Issue', count(*) as 'NbrOfIssues' from (select * from(select a.id as 'UserCDI id', a.intercomExtension, a.personalExtensionId, b.id as 'User id' from UserCDI a left join (select id, userCDIId from User) b on a.id = b.userCDIId)c where c.`User id` is null) as x",
		Query: 		"select * from(select a.id as 'UserCDI id', a.intercomExtension, a.personalExtensionId, b.id as 'User id' from UserCDI a left join (select id, userCDIId from User) b on a.id = b.userCDIId)c where c./`User id/` is null",
		Category: 	"Audit Query", 
		ElId: 		"audit4",		
		Checked: 	true,		
	},
	{
		Issue: 		"Issue#5",	
		IssueName: 	"Users with more than 600 buttons", 
		IssueDef: 	"select 5 as 'Issue#', 'Users with more than 600 buttons' as 'Issue', count(*) as 'NbrOfIssues' from (select distinct c.loginName, c.id as userId, c.userCDIId, b.totalCount as NumOfBtns from User c, Button a inner join (select buttonNumber, parentUserCDIId, count(*) totalCount from Button group by Button.parentUserCDIId having count(*) >=2) b on b.buttonNumber = a.buttonNumber and b.parentUserCDIId = a.parentUserCDIId where c.userCDIId=a.parentUserCDIId and b.totalCount<>600) as x",
		Query: 		"select distinct c.loginName, c.id as userId, c.userCDIId, b.totalCount as NumOfBtns from User c, Button a inner join (select buttonNumber, parentUserCDIId, count(*) totalCount from Button group by Button.parentUserCDIId having count(*) >=2) b on b.buttonNumber = a.buttonNumber and b.parentUserCDIId = a.parentUserCDIId where c.userCDIId=a.parentUserCDIId and b.totalCount<>600",
		Category: 	"Audit Query", 
		ElId: 		"audit5",	
		Checked: 	true,
	},
	{
		Issue: 		"Issue#6",	
		IssueName: 	"Line buttons with no associated line", 
		IssueDef: 	"select 6 as 'Issue#', 'Line buttons with no associated line' as 'Issue', count(*) as 'NbrOfIssues' from (select a.id, a.loginName, b.buttonNumber, b.buttonLabel, b.buttonType, e.resourceAORId, e.appearance, e.resourceAOR from User a, Button b left join (select c.resourceAORId, c.appearance, d.resourceAOR,c.parentButtonId from ButtonResourceAppearance c, ResourceAOR d where d.id=c.resourceAORId) e on b.id=e.parentButtonId where a.userCDIId=b.parentUserCDIId and (b.buttonType='Resource' or b.buttonType='ResourceAndSpeedDial' or b.buttonType='SimplexConference' or b.buttonType='DuplexConference' or b.buttonType='OneButtonDivert') and e.resourceAORId is null order by a.loginName, b.buttonNumber) as x",
		Query: 		"select a.id, a.loginName, b.buttonNumber, b.buttonLabel, b.buttonType, e.resourceAORId, e.appearance, e.resourceAOR from User a, Button b left join (select c.resourceAORId, c.appearance, d.resourceAOR,c.parentButtonId from ButtonResourceAppearance c, ResourceAOR d where d.id=c.resourceAORId) e on b.id=e.parentButtonId where a.userCDIId=b.parentUserCDIId and (b.buttonType='Resource' or b.buttonType='ResourceAndSpeedDial' or b.buttonType='SimplexConference' or b.buttonType='DuplexConference' or b.buttonType='OneButtonDivert') and e.resourceAORId is null order by a.loginName, b.buttonNumber",
		Category: 	"Audit Query", 
		ElId: 		"audit6",	
		Checked: 	true,
	},
	{
		Issue: 		"Issue#7",	
		IssueName: 	"Users with duplicate UserAuthToken (fixed in v5.3sp1)", 
		IssueDef: 	"select 7 as 'Issue#', concat('Users with duplicate UserAuthToken (fixed in v5.3sp1) - ',left(z.softwareVersion,8)) as 'Issue', count(*) as 'NbrOfIssues' from  (select UAT.userId , count(UAT.id) cnt, U.loginName from User U , UserAuthToken  UAT  where U.id=UAT.userId  group by UAT.userId having cnt>1 order by cnt) as x, Zone z where z.id=1",
		Query: 		"select UAT.userId , count(UAT.id) cnt, U.loginName from User U , UserAuthToken  UAT  where U.id=UAT.userId  group by UAT.userId having cnt>1 order by cnt",		
		Category: 	"Audit Query", 
		ElId: 		"audit7",	
		Checked: 	true,
	},
	{
		Issue: 		"Issue#8",	
		IssueName: 	"Large CommHist (record) per day per User", 
		IssueDef: 	"select 8 as 'Issue#', 'Large CommHist (record) per day per User' as 'Issue', count(*) as 'NbrOfIssues' from (select distinct userId, deviceIdId, left(startTime,10) as 'Date', count(startTime) as 'Count' from CommunicationHistory where callType='record' group by left(startTime,10),userId) as x where x.Count>999",
		Query: 		"select distinct userId, deviceIdId, left(startTime,10) as 'Date', count(startTime) as 'Count' from CommunicationHistory where callType='record' group by left(startTime,10),userId",
		Category: 	"Audit Query",
		ElId: 		"audit8",	 
		Checked: 	true,
	},
	{
		Issue: 		"Issue#9",	
		IssueName: 	"Large Directory Tree Structure", 
		IssueDef: 	"select 9 as 'Issue#', 'Large Directory Tree Structure' as 'Issue', count(*) as 'NbrOfIssues' from (select count(*) as 'Count' from DirectoryTreeDirectoryTreeDirectoryTreeSelvesMap) as x where x.Count>250",
		Query: 		"select count(*) as 'Count' from DirectoryTreeDirectoryTreeDirectoryTreeSelvesMap",
		Category: 	"Audit Query", 
		ElId: 		"audit9",
		Checked: 	true,
	},
	{
		Issue: 		"Issue#10",	
		IssueName: 	"Numerous logon loops per day per User", 
		IssueDef: 	"select 10 as 'Issue#', 'Numerous logon loops per day per User' as 'Issue', count(*) as 'NbrOfIssues' from (select distinct left(lh.loginLogoutTime,10) as 'Date', dayname(left(lh.loginLogoutTime,10)) as 'Day', u.loginName, u.id, count(lh.loginLogoutTime) as 'Count' from LoginHistory lh, User u where lh.userId=u.id and lh.loginLogoutType='Login' and lh.clientType='uda_standalone' group by left(lh.loginLogoutTime,10), lh.userId) as x where x.Count>100",
		Query: 		"select distinct left(lh.loginLogoutTime,10) as 'Date', dayname(left(lh.loginLogoutTime,10)) as 'Day', u.loginName, u.id, count(lh.loginLogoutTime) as 'Count' from LoginHistory lh, User u where lh.userId=u.id and lh.loginLogoutType='Login' and lh.clientType='uda_standalone' group by left(lh.loginLogoutTime,10), lh.userId",
		Category: 	"Audit Query", 
		ElId: 		"audit10",	
		Checked: 	true,
	},
	{
		Issue: 		"Issue#11",	
		IssueName: 	"Large database tables", 
		IssueDef: 	"select 11 as 'Issue#', 'Large database tables' as 'Issue', count(*) as 'NbrOfIssues' from (SELECT table_name AS 'Table', (data_length + index_length + data_free) FROM information_schema.TABLES where table_schema=database() and (data_length + index_length + data_free) > 999999999) as x",
		Query: 		"SELECT table_name AS 'Table', (data_length + index_length + data_free) FROM information_schema.TABLES where table_schema=database() and (data_length + index_length + data_free) > 999999999",
		Category: 	"Audit Query", 
		ElId: 		"audit11",	
		Checked: 	true,
	},
	{
		Issue: 		"Issue#12",	
		IssueName: 	"Users with no CLI for RecevedByOthers", 
		IssueDef: 	"select 12 as 'Issue#', 'Users with no CLI for RecevedByOthers' as 'Issue', count(*) as 'NbrOfIssues' from (select u.loginName, uc.cLIPreference from User u, UserCDI uc, (select distinct userId, eventType from CommunicationHistory where eventType='ReceivedByOther' and cLINumber is null group by userId)as z where u.userCDIId=uc.id and uc.cLIPreference='CLI_LOOKUP_NAME' and u.id=z.userId) as x",
		Query: 		"select u.loginName, uc.cLIPreference from User u, UserCDI uc, (select distinct userId, eventType from CommunicationHistory where eventType='ReceivedByOther' and cLINumber is null group by userId)as z where u.userCDIId=uc.id and uc.cLIPreference='CLI_LOOKUP_NAME' and u.id=z.userId",
		Category: 	"Audit Query", 
		ElId: 		"audit12",
		Checked: 	true,
	},
	{
		Issue: 		"Issue#13",	
		IssueName: 	"Bluewave ipcbw account locked", 
		IssueDef: 	"select 13 as 'Issue#', 'Bluewave ipcbw account locked' as 'Issue', count(*) as 'NbrOfIssues' from (select u.loginName, u.accountStatus from User u where u.accountStatus='Locked' and u.loginName='ipcbw') as x",
		Query: 		"select u.loginName, u.accountStatus from User u where u.accountStatus='Locked' and u.loginName='ipcbw'",
		Category: 	"Audit Query", 
		ElId: 		"audit13",
		Checked: 	true,
	},
	{
		Issue: 		"Issue#14",	
		IssueName: 	"USC Prototype with no Location", 
		IssueDef: 	"select 14 as 'Issue#', 'USC Prototype with no Location' as 'Issue', count(*) as 'NbrOfIssues' from (select id, dunkinLocationId, parentZoneId FROM Device where iPAddress=\"1.1.1.4\" and dunkinLocationId is null order by parentZoneId) as x",
		Query: 		"select id, dunkinLocationId, parentZoneId FROM Device where iPAddress=\"1.1.1.4\" and dunkinLocationId is null order by parentZoneId",
		Category: 	"Audit Query", 
		ElId: 		"audit14",
		Checked: 	true,
	},
	{
		Issue: 		"Issue#15",	
		IssueName: 	"Active Speakers with Unequipped lines", 
		IssueDef: 	"select 15 as 'Issue#', 'Active Speakers with Unequipped lines' as 'Issue', count(*) as 'NbrOfIssues' from (SELECT a.loginName, b.speakerNumber, b.label, c.id AS 'resourceAORId', c.resourceAOR, c.equipped, c.lineIsRecorded, g.zoneId AS 'resAorZone', f.zoneId AS 'TrtZone' FROM UserSpeakerChannel b, ResourceAOR c, ResourceZone g, User a LEFT JOIN LogonSession f ON a.id = f.userId WHERE a.userCDIId = b.parentUserCDIId AND b.resourceAORId = c.id AND g.active = TRUE AND a.accountType = 'real' AND c.id = g.parentResourceAORId AND b.activeStatus = 1 AND c.equipped = 0 ORDER BY f.zoneId, a.loginName, b.speakerNumber, g.active DESC, g.zoneId) as x",
		Query: 		"SELECT a.loginName, b.speakerNumber, b.label, c.id AS 'resourceAORId', c.resourceAOR, c.equipped, c.lineIsRecorded, g.zoneId AS 'resAorZone', f.zoneId AS 'TrtZone' FROM UserSpeakerChannel b, ResourceAOR c, ResourceZone g, User a LEFT JOIN LogonSession f ON a.id = f.userId WHERE a.userCDIId = b.parentUserCDIId AND b.resourceAORId = c.id AND g.active = TRUE AND a.accountType = 'real' AND c.id = g.parentResourceAORId AND b.activeStatus = 1 AND c.equipped = 0 ORDER BY f.zoneId, a.loginName, b.speakerNumber, g.active DESC, g.zoneId",
		Category: 	"Audit Query", 
		ElId: 		"audit15",
		Checked: 	true,
	},
}
	*/

const (
	AUDIT_QUERY_1 = `select 1 as 'Issue#', 'Duplicate Buttons' as 'Issue', round(count(*)/2) as 'NbrOfIssues' from (select c.loginName, d.swap, a.*, b.totalCount as Duplicate from User c, UserCDI d, Button a inner join (select buttonNumber, parentUserCDIId, count(*) totalCount from Button group by Button.buttonNumber, Button.parentUserCDIId having count(*) >=2) b on b.buttonNumber = a.buttonNumber and b.parentUserCDIId = a.parentUserCDIId where c.userCDIId=a.parentUserCDIId and c.userCDIId=d.id) as x
union
select 2 as 'Issue#', 'Open Connexion lines on speakers with no listening permissions' as 'Issue', count(*) as 'NbrOfIssues' from (select a.loginName, a.id as userId, a.userCDIId, b.speakerNumber, b.label, b.activeStatus, c.id as resourceAORIdnoLstn from User a, UserSpeakerChannel b, ResourceAOR c where a.userCDIId=b.parentUserCDIId and b.resourceAORId=c.id and c.type='OpenConnexion' and a.accountType='real' and a.id not in (select listenersId from ResourceAORUserListenersMap d where c.id=d.resourceAORId)) as y
union
select 3 as 'Issue#', 'Open Connexion lines on buttons with no listening permissions' as 'Issue', count(*) as 'NbrOfIssues' from (select a.loginName, a.id as userId, a.userCDIId, b.buttonNumber, b.buttonLabel, c.id as resourceAORIdnoListn, e.appearance from User a, Button b, ResourceAOR c, ButtonResourceAppearance e where a.userCDIId=b.parentUserCDIId and e.resourceAORId=c.id and c.type='OpenConnexion' and a.accountType='real' and e.parentButtonId=b.id and a.id not in (select listenersId from ResourceAORUserListenersMap d where c.id=d.resourceAORId)) as x
union
select 4 as 'Issue#', 'UserCDI with no User id' as 'Issue', count(*) as 'NbrOfIssues' from (select * from(select a.id as 'UserCDI id', a.intercomExtension, a.personalExtensionId, b.id as 'User id' from UserCDI a left join (select id, userCDIId from User) b on a.id = b.userCDIId)c where c.'User id' is null) as x
union
select 5 as 'Issue#', 'Users with more than 600 buttons' as 'Issue', count(*) as 'NbrOfIssues' from (select distinct c.loginName, c.id as userId, c.userCDIId, b.totalCount as NumOfBtns from User c, Button a inner join (select buttonNumber, parentUserCDIId, count(*) totalCount from Button group by Button.parentUserCDIId having count(*) >=2) b on b.buttonNumber = a.buttonNumber and b.parentUserCDIId = a.parentUserCDIId where c.userCDIId=a.parentUserCDIId and b.totalCount<>600) as x
union
select 6 as 'Issue#', 'Line buttons with no associated line' as 'Issue', count(*) as 'NbrOfIssues' from (select a.id, a.loginName, b.buttonNumber, b.buttonLabel, b.buttonType, e.resourceAORId, e.appearance, e.resourceAOR from User a, Button b left join (select c.resourceAORId, c.appearance, d.resourceAOR,c.parentButtonId from ButtonResourceAppearance c, ResourceAOR d where d.id=c.resourceAORId) e on b.id=e.parentButtonId where a.userCDIId=b.parentUserCDIId and (b.buttonType='Resource' or b.buttonType='ResourceAndSpeedDial' or b.buttonType='SimplexConference' or b.buttonType='DuplexConference' or b.buttonType='OneButtonDivert') and e.resourceAORId is null order by a.loginName, b.buttonNumber) as x
union
select 7 as 'Issue#', concat('Users with duplicate UserAuthToken (fixed in v5.3sp1) - ',left(z.softwareVersion,8)) as 'Issue', count(*) as 'NbrOfIssues' from  (select UAT.userId , count(UAT.id) cnt, U.loginName from User U , UserAuthToken  UAT  where U.id=UAT.userId  group by UAT.userId having cnt>1 order by cnt) as x, Zone z where z.id=1
union
select 8 as 'Issue#', 'Large CommHist (record) per day per User' as 'Issue', count(*) as 'NbrOfIssues' from (select distinct userId, deviceIdId, left(startTime,10) as 'Date', count(startTime) as 'Count' from CommunicationHistory where callType='record' group by left(startTime,10),userId) as x where x.Count>999
union
select 9 as 'Issue#', 'Large Directory Tree Structure' as 'Issue', count(*) as 'NbrOfIssues' from (select count(*) as 'Count' from DirectoryTreeDirectoryTreeDirectoryTreeSelvesMap) as x where x.Count>250
union
select 10 as 'Issue#', 'Numerous logon loops per day per User' as 'Issue', count(*) as 'NbrOfIssues' from (select distinct left(lh.loginLogoutTime,10) as 'Date', dayname(left(lh.loginLogoutTime,10)) as 'Day', u.loginName, u.id, count(lh.loginLogoutTime) as 'Count' from LoginHistory lh, User u where lh.userId=u.id and lh.loginLogoutType='Login' and lh.clientType='uda_standalone' group by left(lh.loginLogoutTime,10), lh.userId) as x where x.Count>100
union
select 11 as 'Issue#', 'Large database tables' as 'Issue', count(*) as 'NbrOfIssues' from (SELECT table_name AS 'Table', (data_length + index_length + data_free) FROM information_schema.TABLES where table_schema=database() and (data_length + index_length + data_free) > 999999999) as x
union
select 12 as 'Issue#', 'Users with no CLI for RecevedByOthers' as 'Issue', count(*) as 'NbrOfIssues' from (select u.loginName, uc.cLIPreference from User u, UserCDI uc, (select distinct userId, eventType from CommunicationHistory where eventType='ReceivedByOther' and cLINumber is null group by userId)as z where u.userCDIId=uc.id and uc.cLIPreference='CLI_LOOKUP_NAME' and u.id=z.userId) as x
union
select 13 as 'Issue#', 'Bluewave ipcbw account locked' as 'Issue', count(*) as 'NbrOfIssues' from (select u.loginName, u.accountStatus from User u where u.accountStatus='Locked' and u.loginName='ipcbw') as x`
	AUDIT_QUERY_2 = ``
)