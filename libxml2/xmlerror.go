package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type ErrorLevel c.Int

const (
	ERRNONE    ErrorLevel = 0
	ERRWARNING ErrorLevel = 1
	ERRERROR   ErrorLevel = 2
	ERRFATAL   ErrorLevel = 3
)

type ErrorDomain c.Int

const (
	FROMNONE        ErrorDomain = 0
	FROMPARSER      ErrorDomain = 1
	FROMTREE        ErrorDomain = 2
	FROMNAMESPACE   ErrorDomain = 3
	FROMDTD         ErrorDomain = 4
	FROMHTML        ErrorDomain = 5
	FROMMEMORY      ErrorDomain = 6
	FROMOUTPUT      ErrorDomain = 7
	FROMIO          ErrorDomain = 8
	FROMFTP         ErrorDomain = 9
	FROMHTTP        ErrorDomain = 10
	FROMXINCLUDE    ErrorDomain = 11
	FROMXPATH       ErrorDomain = 12
	FROMXPOINTER    ErrorDomain = 13
	FROMREGEXP      ErrorDomain = 14
	FROMDATATYPE    ErrorDomain = 15
	FROMSCHEMASP    ErrorDomain = 16
	FROMSCHEMASV    ErrorDomain = 17
	FROMRELAXNGP    ErrorDomain = 18
	FROMRELAXNGV    ErrorDomain = 19
	FROMCATALOG     ErrorDomain = 20
	FROMC14N        ErrorDomain = 21
	FROMXSLT        ErrorDomain = 22
	FROMVALID       ErrorDomain = 23
	FROMCHECK       ErrorDomain = 24
	FROMWRITER      ErrorDomain = 25
	FROMMODULE      ErrorDomain = 26
	FROMI18N        ErrorDomain = 27
	FROMSCHEMATRONV ErrorDomain = 28
	FROMBUFFER      ErrorDomain = 29
	FROMURI         ErrorDomain = 30
)

type X_xmlError struct {
	Domain  c.Int
	Code    c.Int
	Message *int8
	Level   ErrorLevel
	File    *int8
	Line    c.Int
	Str1    *int8
	Str2    *int8
	Str3    *int8
	Int1    c.Int
	Int2    c.Int
	Ctxt    unsafe.Pointer
	Node    unsafe.Pointer
}
type Error X_xmlError
type ErrorPtr *Error
type ParserErrors c.Int

const (
	ERROK                                   ParserErrors = 0
	ERRINTERNALERROR                        ParserErrors = 1
	ERRNOMEMORY                             ParserErrors = 2
	ERRDOCUMENTSTART                        ParserErrors = 3
	ERRDOCUMENTEMPTY                        ParserErrors = 4
	ERRDOCUMENTEND                          ParserErrors = 5
	ERRINVALIDHEXCHARREF                    ParserErrors = 6
	ERRINVALIDDECCHARREF                    ParserErrors = 7
	ERRINVALIDCHARREF                       ParserErrors = 8
	ERRINVALIDCHAR                          ParserErrors = 9
	ERRCHARREFATEOF                         ParserErrors = 10
	ERRCHARREFINPROLOG                      ParserErrors = 11
	ERRCHARREFINEPILOG                      ParserErrors = 12
	ERRCHARREFINDTD                         ParserErrors = 13
	ERRENTITYREFATEOF                       ParserErrors = 14
	ERRENTITYREFINPROLOG                    ParserErrors = 15
	ERRENTITYREFINEPILOG                    ParserErrors = 16
	ERRENTITYREFINDTD                       ParserErrors = 17
	ERRPEREFATEOF                           ParserErrors = 18
	ERRPEREFINPROLOG                        ParserErrors = 19
	ERRPEREFINEPILOG                        ParserErrors = 20
	ERRPEREFININTSUBSET                     ParserErrors = 21
	ERRENTITYREFNONAME                      ParserErrors = 22
	ERRENTITYREFSEMICOLMISSING              ParserErrors = 23
	ERRPEREFNONAME                          ParserErrors = 24
	ERRPEREFSEMICOLMISSING                  ParserErrors = 25
	ERRUNDECLAREDENTITY                     ParserErrors = 26
	WARUNDECLAREDENTITY                     ParserErrors = 27
	ERRUNPARSEDENTITY                       ParserErrors = 28
	ERRENTITYISEXTERNAL                     ParserErrors = 29
	ERRENTITYISPARAMETER                    ParserErrors = 30
	ERRUNKNOWNENCODING                      ParserErrors = 31
	ERRUNSUPPORTEDENCODING                  ParserErrors = 32
	ERRSTRINGNOTSTARTED                     ParserErrors = 33
	ERRSTRINGNOTCLOSED                      ParserErrors = 34
	ERRNSDECLERROR                          ParserErrors = 35
	ERRENTITYNOTSTARTED                     ParserErrors = 36
	ERRENTITYNOTFINISHED                    ParserErrors = 37
	ERRLTINATTRIBUTE                        ParserErrors = 38
	ERRATTRIBUTENOTSTARTED                  ParserErrors = 39
	ERRATTRIBUTENOTFINISHED                 ParserErrors = 40
	ERRATTRIBUTEWITHOUTVALUE                ParserErrors = 41
	ERRATTRIBUTEREDEFINED                   ParserErrors = 42
	ERRLITERALNOTSTARTED                    ParserErrors = 43
	ERRLITERALNOTFINISHED                   ParserErrors = 44
	ERRCOMMENTNOTFINISHED                   ParserErrors = 45
	ERRPINOTSTARTED                         ParserErrors = 46
	ERRPINOTFINISHED                        ParserErrors = 47
	ERRNOTATIONNOTSTARTED                   ParserErrors = 48
	ERRNOTATIONNOTFINISHED                  ParserErrors = 49
	ERRATTLISTNOTSTARTED                    ParserErrors = 50
	ERRATTLISTNOTFINISHED                   ParserErrors = 51
	ERRMIXEDNOTSTARTED                      ParserErrors = 52
	ERRMIXEDNOTFINISHED                     ParserErrors = 53
	ERRELEMCONTENTNOTSTARTED                ParserErrors = 54
	ERRELEMCONTENTNOTFINISHED               ParserErrors = 55
	ERRXMLDECLNOTSTARTED                    ParserErrors = 56
	ERRXMLDECLNOTFINISHED                   ParserErrors = 57
	ERRCONDSECNOTSTARTED                    ParserErrors = 58
	ERRCONDSECNOTFINISHED                   ParserErrors = 59
	ERREXTSUBSETNOTFINISHED                 ParserErrors = 60
	ERRDOCTYPENOTFINISHED                   ParserErrors = 61
	ERRMISPLACEDCDATAEND                    ParserErrors = 62
	ERRCDATANOTFINISHED                     ParserErrors = 63
	ERRRESERVEDXMLNAME                      ParserErrors = 64
	ERRSPACEREQUIRED                        ParserErrors = 65
	ERRSEPARATORREQUIRED                    ParserErrors = 66
	ERRNMTOKENREQUIRED                      ParserErrors = 67
	ERRNAMEREQUIRED                         ParserErrors = 68
	ERRPCDATAREQUIRED                       ParserErrors = 69
	ERRURIREQUIRED                          ParserErrors = 70
	ERRPUBIDREQUIRED                        ParserErrors = 71
	ERRLTREQUIRED                           ParserErrors = 72
	ERRGTREQUIRED                           ParserErrors = 73
	ERRLTSLASHREQUIRED                      ParserErrors = 74
	ERREQUALREQUIRED                        ParserErrors = 75
	ERRTAGNAMEMISMATCH                      ParserErrors = 76
	ERRTAGNOTFINISHED                       ParserErrors = 77
	ERRSTANDALONEVALUE                      ParserErrors = 78
	ERRENCODINGNAME                         ParserErrors = 79
	ERRHYPHENINCOMMENT                      ParserErrors = 80
	ERRINVALIDENCODING                      ParserErrors = 81
	ERREXTENTITYSTANDALONE                  ParserErrors = 82
	ERRCONDSECINVALID                       ParserErrors = 83
	ERRVALUEREQUIRED                        ParserErrors = 84
	ERRNOTWELLBALANCED                      ParserErrors = 85
	ERREXTRACONTENT                         ParserErrors = 86
	ERRENTITYCHARERROR                      ParserErrors = 87
	ERRENTITYPEINTERNAL                     ParserErrors = 88
	ERRENTITYLOOP                           ParserErrors = 89
	ERRENTITYBOUNDARY                       ParserErrors = 90
	ERRINVALIDURI                           ParserErrors = 91
	ERRURIFRAGMENT                          ParserErrors = 92
	WARCATALOGPI                            ParserErrors = 93
	ERRNODTD                                ParserErrors = 94
	ERRCONDSECINVALIDKEYWORD                ParserErrors = 95
	ERRVERSIONMISSING                       ParserErrors = 96
	WARUNKNOWNVERSION                       ParserErrors = 97
	WARLANGVALUE                            ParserErrors = 98
	WARNSURI                                ParserErrors = 99
	WARNSURIRELATIVE                        ParserErrors = 100
	ERRMISSINGENCODING                      ParserErrors = 101
	WARSPACEVALUE                           ParserErrors = 102
	ERRNOTSTANDALONE                        ParserErrors = 103
	ERRENTITYPROCESSING                     ParserErrors = 104
	ERRNOTATIONPROCESSING                   ParserErrors = 105
	WARNSCOLUMN                             ParserErrors = 106
	WARENTITYREDEFINED                      ParserErrors = 107
	ERRUNKNOWNVERSION                       ParserErrors = 108
	ERRVERSIONMISMATCH                      ParserErrors = 109
	ERRNAMETOOLONG                          ParserErrors = 110
	ERRUSERSTOP                             ParserErrors = 111
	ERRCOMMENTABRUPTLYENDED                 ParserErrors = 112
	WARENCODINGMISMATCH                     ParserErrors = 113
	ERRRESOURCELIMIT                        ParserErrors = 114
	ERRARGUMENT                             ParserErrors = 115
	ERRSYSTEM                               ParserErrors = 116
	ERRREDECLPREDEFENTITY                   ParserErrors = 117
	ERRINTSUBSETNOTFINISHED                 ParserErrors = 118
	NSERRXMLNAMESPACE                       ParserErrors = 200
	NSERRUNDEFINEDNAMESPACE                 ParserErrors = 201
	NSERRQNAME                              ParserErrors = 202
	NSERRATTRIBUTEREDEFINED                 ParserErrors = 203
	NSERREMPTY                              ParserErrors = 204
	NSERRCOLON                              ParserErrors = 205
	DTDATTRIBUTEDEFAULT                     ParserErrors = 500
	DTDATTRIBUTEREDEFINED                   ParserErrors = 501
	DTDATTRIBUTEVALUE                       ParserErrors = 502
	DTDCONTENTERROR                         ParserErrors = 503
	DTDCONTENTMODEL                         ParserErrors = 504
	DTDCONTENTNOTDETERMINIST                ParserErrors = 505
	DTDDIFFERENTPREFIX                      ParserErrors = 506
	DTDELEMDEFAULTNAMESPACE                 ParserErrors = 507
	DTDELEMNAMESPACE                        ParserErrors = 508
	DTDELEMREDEFINED                        ParserErrors = 509
	DTDEMPTYNOTATION                        ParserErrors = 510
	DTDENTITYTYPE                           ParserErrors = 511
	DTDIDFIXED                              ParserErrors = 512
	DTDIDREDEFINED                          ParserErrors = 513
	DTDIDSUBSET                             ParserErrors = 514
	DTDINVALIDCHILD                         ParserErrors = 515
	DTDINVALIDDEFAULT                       ParserErrors = 516
	DTDLOADERROR                            ParserErrors = 517
	DTDMISSINGATTRIBUTE                     ParserErrors = 518
	DTDMIXEDCORRUPT                         ParserErrors = 519
	DTDMULTIPLEID                           ParserErrors = 520
	DTDNODOC                                ParserErrors = 521
	DTDNODTD                                ParserErrors = 522
	DTDNOELEMNAME                           ParserErrors = 523
	DTDNOPREFIX                             ParserErrors = 524
	DTDNOROOT                               ParserErrors = 525
	DTDNOTATIONREDEFINED                    ParserErrors = 526
	DTDNOTATIONVALUE                        ParserErrors = 527
	DTDNOTEMPTY                             ParserErrors = 528
	DTDNOTPCDATA                            ParserErrors = 529
	DTDNOTSTANDALONE                        ParserErrors = 530
	DTDROOTNAME                             ParserErrors = 531
	DTDSTANDALONEWHITESPACE                 ParserErrors = 532
	DTDUNKNOWNATTRIBUTE                     ParserErrors = 533
	DTDUNKNOWNELEM                          ParserErrors = 534
	DTDUNKNOWNENTITY                        ParserErrors = 535
	DTDUNKNOWNID                            ParserErrors = 536
	DTDUNKNOWNNOTATION                      ParserErrors = 537
	DTDSTANDALONEDEFAULTED                  ParserErrors = 538
	DTDXMLIDVALUE                           ParserErrors = 539
	DTDXMLIDTYPE                            ParserErrors = 540
	DTDDUPTOKEN                             ParserErrors = 541
	HTMLSTRUCUREERROR                       ParserErrors = 800
	HTMLUNKNOWNTAG                          ParserErrors = 801
	HTMLINCORRECTLYOPENEDCOMMENT            ParserErrors = 802
	RNGPANYNAMEATTRANCESTOR                 ParserErrors = 1000
	RNGPATTRCONFLICT                        ParserErrors = 1001
	RNGPATTRIBUTECHILDREN                   ParserErrors = 1002
	RNGPATTRIBUTECONTENT                    ParserErrors = 1003
	RNGPATTRIBUTEEMPTY                      ParserErrors = 1004
	RNGPATTRIBUTENOOP                       ParserErrors = 1005
	RNGPCHOICECONTENT                       ParserErrors = 1006
	RNGPCHOICEEMPTY                         ParserErrors = 1007
	RNGPCREATEFAILURE                       ParserErrors = 1008
	RNGPDATACONTENT                         ParserErrors = 1009
	RNGPDEFCHOICEANDINTERLEAVE              ParserErrors = 1010
	RNGPDEFINECREATEFAILED                  ParserErrors = 1011
	RNGPDEFINEEMPTY                         ParserErrors = 1012
	RNGPDEFINEMISSING                       ParserErrors = 1013
	RNGPDEFINENAMEMISSING                   ParserErrors = 1014
	RNGPELEMCONTENTEMPTY                    ParserErrors = 1015
	RNGPELEMCONTENTERROR                    ParserErrors = 1016
	RNGPELEMENTEMPTY                        ParserErrors = 1017
	RNGPELEMENTCONTENT                      ParserErrors = 1018
	RNGPELEMENTNAME                         ParserErrors = 1019
	RNGPELEMENTNOCONTENT                    ParserErrors = 1020
	RNGPELEMTEXTCONFLICT                    ParserErrors = 1021
	RNGPEMPTY                               ParserErrors = 1022
	RNGPEMPTYCONSTRUCT                      ParserErrors = 1023
	RNGPEMPTYCONTENT                        ParserErrors = 1024
	RNGPEMPTYNOTEMPTY                       ParserErrors = 1025
	RNGPERRORTYPELIB                        ParserErrors = 1026
	RNGPEXCEPTEMPTY                         ParserErrors = 1027
	RNGPEXCEPTMISSING                       ParserErrors = 1028
	RNGPEXCEPTMULTIPLE                      ParserErrors = 1029
	RNGPEXCEPTNOCONTENT                     ParserErrors = 1030
	RNGPEXTERNALREFEMTPY                    ParserErrors = 1031
	RNGPEXTERNALREFFAILURE                  ParserErrors = 1032
	RNGPEXTERNALREFRECURSE                  ParserErrors = 1033
	RNGPFORBIDDENATTRIBUTE                  ParserErrors = 1034
	RNGPFOREIGNELEMENT                      ParserErrors = 1035
	RNGPGRAMMARCONTENT                      ParserErrors = 1036
	RNGPGRAMMAREMPTY                        ParserErrors = 1037
	RNGPGRAMMARMISSING                      ParserErrors = 1038
	RNGPGRAMMARNOSTART                      ParserErrors = 1039
	RNGPGROUPATTRCONFLICT                   ParserErrors = 1040
	RNGPHREFERROR                           ParserErrors = 1041
	RNGPINCLUDEEMPTY                        ParserErrors = 1042
	RNGPINCLUDEFAILURE                      ParserErrors = 1043
	RNGPINCLUDERECURSE                      ParserErrors = 1044
	RNGPINTERLEAVEADD                       ParserErrors = 1045
	RNGPINTERLEAVECREATEFAILED              ParserErrors = 1046
	RNGPINTERLEAVEEMPTY                     ParserErrors = 1047
	RNGPINTERLEAVENOCONTENT                 ParserErrors = 1048
	RNGPINVALIDDEFINENAME                   ParserErrors = 1049
	RNGPINVALIDURI                          ParserErrors = 1050
	RNGPINVALIDVALUE                        ParserErrors = 1051
	RNGPMISSINGHREF                         ParserErrors = 1052
	RNGPNAMEMISSING                         ParserErrors = 1053
	RNGPNEEDCOMBINE                         ParserErrors = 1054
	RNGPNOTALLOWEDNOTEMPTY                  ParserErrors = 1055
	RNGPNSNAMEATTRANCESTOR                  ParserErrors = 1056
	RNGPNSNAMENONS                          ParserErrors = 1057
	RNGPPARAMFORBIDDEN                      ParserErrors = 1058
	RNGPPARAMNAMEMISSING                    ParserErrors = 1059
	RNGPPARENTREFCREATEFAILED               ParserErrors = 1060
	RNGPPARENTREFNAMEINVALID                ParserErrors = 1061
	RNGPPARENTREFNONAME                     ParserErrors = 1062
	RNGPPARENTREFNOPARENT                   ParserErrors = 1063
	RNGPPARENTREFNOTEMPTY                   ParserErrors = 1064
	RNGPPARSEERROR                          ParserErrors = 1065
	RNGPPATANYNAMEEXCEPTANYNAME             ParserErrors = 1066
	RNGPPATATTRATTR                         ParserErrors = 1067
	RNGPPATATTRELEM                         ParserErrors = 1068
	RNGPPATDATAEXCEPTATTR                   ParserErrors = 1069
	RNGPPATDATAEXCEPTELEM                   ParserErrors = 1070
	RNGPPATDATAEXCEPTEMPTY                  ParserErrors = 1071
	RNGPPATDATAEXCEPTGROUP                  ParserErrors = 1072
	RNGPPATDATAEXCEPTINTERLEAVE             ParserErrors = 1073
	RNGPPATDATAEXCEPTLIST                   ParserErrors = 1074
	RNGPPATDATAEXCEPTONEMORE                ParserErrors = 1075
	RNGPPATDATAEXCEPTREF                    ParserErrors = 1076
	RNGPPATDATAEXCEPTTEXT                   ParserErrors = 1077
	RNGPPATLISTATTR                         ParserErrors = 1078
	RNGPPATLISTELEM                         ParserErrors = 1079
	RNGPPATLISTINTERLEAVE                   ParserErrors = 1080
	RNGPPATLISTLIST                         ParserErrors = 1081
	RNGPPATLISTREF                          ParserErrors = 1082
	RNGPPATLISTTEXT                         ParserErrors = 1083
	RNGPPATNSNAMEEXCEPTANYNAME              ParserErrors = 1084
	RNGPPATNSNAMEEXCEPTNSNAME               ParserErrors = 1085
	RNGPPATONEMOREGROUPATTR                 ParserErrors = 1086
	RNGPPATONEMOREINTERLEAVEATTR            ParserErrors = 1087
	RNGPPATSTARTATTR                        ParserErrors = 1088
	RNGPPATSTARTDATA                        ParserErrors = 1089
	RNGPPATSTARTEMPTY                       ParserErrors = 1090
	RNGPPATSTARTGROUP                       ParserErrors = 1091
	RNGPPATSTARTINTERLEAVE                  ParserErrors = 1092
	RNGPPATSTARTLIST                        ParserErrors = 1093
	RNGPPATSTARTONEMORE                     ParserErrors = 1094
	RNGPPATSTARTTEXT                        ParserErrors = 1095
	RNGPPATSTARTVALUE                       ParserErrors = 1096
	RNGPPREFIXUNDEFINED                     ParserErrors = 1097
	RNGPREFCREATEFAILED                     ParserErrors = 1098
	RNGPREFCYCLE                            ParserErrors = 1099
	RNGPREFNAMEINVALID                      ParserErrors = 1100
	RNGPREFNODEF                            ParserErrors = 1101
	RNGPREFNONAME                           ParserErrors = 1102
	RNGPREFNOTEMPTY                         ParserErrors = 1103
	RNGPSTARTCHOICEANDINTERLEAVE            ParserErrors = 1104
	RNGPSTARTCONTENT                        ParserErrors = 1105
	RNGPSTARTEMPTY                          ParserErrors = 1106
	RNGPSTARTMISSING                        ParserErrors = 1107
	RNGPTEXTEXPECTED                        ParserErrors = 1108
	RNGPTEXTHASCHILD                        ParserErrors = 1109
	RNGPTYPEMISSING                         ParserErrors = 1110
	RNGPTYPENOTFOUND                        ParserErrors = 1111
	RNGPTYPEVALUE                           ParserErrors = 1112
	RNGPUNKNOWNATTRIBUTE                    ParserErrors = 1113
	RNGPUNKNOWNCOMBINE                      ParserErrors = 1114
	RNGPUNKNOWNCONSTRUCT                    ParserErrors = 1115
	RNGPUNKNOWNTYPELIB                      ParserErrors = 1116
	RNGPURIFRAGMENT                         ParserErrors = 1117
	RNGPURINOTABSOLUTE                      ParserErrors = 1118
	RNGPVALUEEMPTY                          ParserErrors = 1119
	RNGPVALUENOCONTENT                      ParserErrors = 1120
	RNGPXMLNSNAME                           ParserErrors = 1121
	RNGPXMLNS                               ParserErrors = 1122
	XPATHEXPRESSIONOK                       ParserErrors = 1200
	XPATHNUMBERERROR                        ParserErrors = 1201
	XPATHUNFINISHEDLITERALERROR             ParserErrors = 1202
	XPATHSTARTLITERALERROR                  ParserErrors = 1203
	XPATHVARIABLEREFERROR                   ParserErrors = 1204
	XPATHUNDEFVARIABLEERROR                 ParserErrors = 1205
	XPATHINVALIDPREDICATEERROR              ParserErrors = 1206
	XPATHEXPRERROR                          ParserErrors = 1207
	XPATHUNCLOSEDERROR                      ParserErrors = 1208
	XPATHUNKNOWNFUNCERROR                   ParserErrors = 1209
	XPATHINVALIDOPERAND                     ParserErrors = 1210
	XPATHINVALIDTYPE                        ParserErrors = 1211
	XPATHINVALIDARITY                       ParserErrors = 1212
	XPATHINVALIDCTXTSIZE                    ParserErrors = 1213
	XPATHINVALIDCTXTPOSITION                ParserErrors = 1214
	XPATHMEMORYERROR                        ParserErrors = 1215
	XPTRSYNTAXERROR                         ParserErrors = 1216
	XPTRRESOURCEERROR                       ParserErrors = 1217
	XPTRSUBRESOURCEERROR                    ParserErrors = 1218
	XPATHUNDEFPREFIXERROR                   ParserErrors = 1219
	XPATHENCODINGERROR                      ParserErrors = 1220
	XPATHINVALIDCHARERROR                   ParserErrors = 1221
	TREEINVALIDHEX                          ParserErrors = 1300
	TREEINVALIDDEC                          ParserErrors = 1301
	TREEUNTERMINATEDENTITY                  ParserErrors = 1302
	TREENOTUTF8                             ParserErrors = 1303
	SAVENOTUTF8                             ParserErrors = 1400
	SAVECHARINVALID                         ParserErrors = 1401
	SAVENODOCTYPE                           ParserErrors = 1402
	SAVEUNKNOWNENCODING                     ParserErrors = 1403
	REGEXPCOMPILEERROR                      ParserErrors = 1450
	IOUNKNOWN                               ParserErrors = 1500
	IOEACCES                                ParserErrors = 1501
	IOEAGAIN                                ParserErrors = 1502
	IOEBADF                                 ParserErrors = 1503
	IOEBADMSG                               ParserErrors = 1504
	IOEBUSY                                 ParserErrors = 1505
	IOECANCELED                             ParserErrors = 1506
	IOECHILD                                ParserErrors = 1507
	IOEDEADLK                               ParserErrors = 1508
	IOEDOM                                  ParserErrors = 1509
	IOEEXIST                                ParserErrors = 1510
	IOEFAULT                                ParserErrors = 1511
	IOEFBIG                                 ParserErrors = 1512
	IOEINPROGRESS                           ParserErrors = 1513
	IOEINTR                                 ParserErrors = 1514
	IOEINVAL                                ParserErrors = 1515
	IOEIO                                   ParserErrors = 1516
	IOEISDIR                                ParserErrors = 1517
	IOEMFILE                                ParserErrors = 1518
	IOEMLINK                                ParserErrors = 1519
	IOEMSGSIZE                              ParserErrors = 1520
	IOENAMETOOLONG                          ParserErrors = 1521
	IOENFILE                                ParserErrors = 1522
	IOENODEV                                ParserErrors = 1523
	IOENOENT                                ParserErrors = 1524
	IOENOEXEC                               ParserErrors = 1525
	IOENOLCK                                ParserErrors = 1526
	IOENOMEM                                ParserErrors = 1527
	IOENOSPC                                ParserErrors = 1528
	IOENOSYS                                ParserErrors = 1529
	IOENOTDIR                               ParserErrors = 1530
	IOENOTEMPTY                             ParserErrors = 1531
	IOENOTSUP                               ParserErrors = 1532
	IOENOTTY                                ParserErrors = 1533
	IOENXIO                                 ParserErrors = 1534
	IOEPERM                                 ParserErrors = 1535
	IOEPIPE                                 ParserErrors = 1536
	IOERANGE                                ParserErrors = 1537
	IOEROFS                                 ParserErrors = 1538
	IOESPIPE                                ParserErrors = 1539
	IOESRCH                                 ParserErrors = 1540
	IOETIMEDOUT                             ParserErrors = 1541
	IOEXDEV                                 ParserErrors = 1542
	IONETWORKATTEMPT                        ParserErrors = 1543
	IOENCODER                               ParserErrors = 1544
	IOFLUSH                                 ParserErrors = 1545
	IOWRITE                                 ParserErrors = 1546
	IONOINPUT                               ParserErrors = 1547
	IOBUFFERFULL                            ParserErrors = 1548
	IOLOADERROR                             ParserErrors = 1549
	IOENOTSOCK                              ParserErrors = 1550
	IOEISCONN                               ParserErrors = 1551
	IOECONNREFUSED                          ParserErrors = 1552
	IOENETUNREACH                           ParserErrors = 1553
	IOEADDRINUSE                            ParserErrors = 1554
	IOEALREADY                              ParserErrors = 1555
	IOEAFNOSUPPORT                          ParserErrors = 1556
	IOUNSUPPORTEDPROTOCOL                   ParserErrors = 1557
	XINCLUDERECURSION                       ParserErrors = 1600
	XINCLUDEPARSEVALUE                      ParserErrors = 1601
	XINCLUDEENTITYDEFMISMATCH               ParserErrors = 1602
	XINCLUDENOHREF                          ParserErrors = 1603
	XINCLUDENOFALLBACK                      ParserErrors = 1604
	XINCLUDEHREFURI                         ParserErrors = 1605
	XINCLUDETEXTFRAGMENT                    ParserErrors = 1606
	XINCLUDETEXTDOCUMENT                    ParserErrors = 1607
	XINCLUDEINVALIDCHAR                     ParserErrors = 1608
	XINCLUDEBUILDFAILED                     ParserErrors = 1609
	XINCLUDEUNKNOWNENCODING                 ParserErrors = 1610
	XINCLUDEMULTIPLEROOT                    ParserErrors = 1611
	XINCLUDEXPTRFAILED                      ParserErrors = 1612
	XINCLUDEXPTRRESULT                      ParserErrors = 1613
	XINCLUDEINCLUDEININCLUDE                ParserErrors = 1614
	XINCLUDEFALLBACKSININCLUDE              ParserErrors = 1615
	XINCLUDEFALLBACKNOTININCLUDE            ParserErrors = 1616
	XINCLUDEDEPRECATEDNS                    ParserErrors = 1617
	XINCLUDEFRAGMENTID                      ParserErrors = 1618
	CATALOGMISSINGATTR                      ParserErrors = 1650
	CATALOGENTRYBROKEN                      ParserErrors = 1651
	CATALOGPREFERVALUE                      ParserErrors = 1652
	CATALOGNOTCATALOG                       ParserErrors = 1653
	CATALOGRECURSION                        ParserErrors = 1654
	SCHEMAPPREFIXUNDEFINED                  ParserErrors = 1700
	SCHEMAPATTRFORMDEFAULTVALUE             ParserErrors = 1701
	SCHEMAPATTRGRPNONAMENOREF               ParserErrors = 1702
	SCHEMAPATTRNONAMENOREF                  ParserErrors = 1703
	SCHEMAPCOMPLEXTYPENONAMENOREF           ParserErrors = 1704
	SCHEMAPELEMFORMDEFAULTVALUE             ParserErrors = 1705
	SCHEMAPELEMNONAMENOREF                  ParserErrors = 1706
	SCHEMAPEXTENSIONNOBASE                  ParserErrors = 1707
	SCHEMAPFACETNOVALUE                     ParserErrors = 1708
	SCHEMAPFAILEDBUILDIMPORT                ParserErrors = 1709
	SCHEMAPGROUPNONAMENOREF                 ParserErrors = 1710
	SCHEMAPIMPORTNAMESPACENOTURI            ParserErrors = 1711
	SCHEMAPIMPORTREDEFINENSNAME             ParserErrors = 1712
	SCHEMAPIMPORTSCHEMANOTURI               ParserErrors = 1713
	SCHEMAPINVALIDBOOLEAN                   ParserErrors = 1714
	SCHEMAPINVALIDENUM                      ParserErrors = 1715
	SCHEMAPINVALIDFACET                     ParserErrors = 1716
	SCHEMAPINVALIDFACETVALUE                ParserErrors = 1717
	SCHEMAPINVALIDMAXOCCURS                 ParserErrors = 1718
	SCHEMAPINVALIDMINOCCURS                 ParserErrors = 1719
	SCHEMAPINVALIDREFANDSUBTYPE             ParserErrors = 1720
	SCHEMAPINVALIDWHITESPACE                ParserErrors = 1721
	SCHEMAPNOATTRNOREF                      ParserErrors = 1722
	SCHEMAPNOTATIONNONAME                   ParserErrors = 1723
	SCHEMAPNOTYPENOREF                      ParserErrors = 1724
	SCHEMAPREFANDSUBTYPE                    ParserErrors = 1725
	SCHEMAPRESTRICTIONNONAMENOREF           ParserErrors = 1726
	SCHEMAPSIMPLETYPENONAME                 ParserErrors = 1727
	SCHEMAPTYPEANDSUBTYPE                   ParserErrors = 1728
	SCHEMAPUNKNOWNALLCHILD                  ParserErrors = 1729
	SCHEMAPUNKNOWNANYATTRIBUTECHILD         ParserErrors = 1730
	SCHEMAPUNKNOWNATTRCHILD                 ParserErrors = 1731
	SCHEMAPUNKNOWNATTRGRPCHILD              ParserErrors = 1732
	SCHEMAPUNKNOWNATTRIBUTEGROUP            ParserErrors = 1733
	SCHEMAPUNKNOWNBASETYPE                  ParserErrors = 1734
	SCHEMAPUNKNOWNCHOICECHILD               ParserErrors = 1735
	SCHEMAPUNKNOWNCOMPLEXCONTENTCHILD       ParserErrors = 1736
	SCHEMAPUNKNOWNCOMPLEXTYPECHILD          ParserErrors = 1737
	SCHEMAPUNKNOWNELEMCHILD                 ParserErrors = 1738
	SCHEMAPUNKNOWNEXTENSIONCHILD            ParserErrors = 1739
	SCHEMAPUNKNOWNFACETCHILD                ParserErrors = 1740
	SCHEMAPUNKNOWNFACETTYPE                 ParserErrors = 1741
	SCHEMAPUNKNOWNGROUPCHILD                ParserErrors = 1742
	SCHEMAPUNKNOWNIMPORTCHILD               ParserErrors = 1743
	SCHEMAPUNKNOWNLISTCHILD                 ParserErrors = 1744
	SCHEMAPUNKNOWNNOTATIONCHILD             ParserErrors = 1745
	SCHEMAPUNKNOWNPROCESSCONTENTCHILD       ParserErrors = 1746
	SCHEMAPUNKNOWNREF                       ParserErrors = 1747
	SCHEMAPUNKNOWNRESTRICTIONCHILD          ParserErrors = 1748
	SCHEMAPUNKNOWNSCHEMASCHILD              ParserErrors = 1749
	SCHEMAPUNKNOWNSEQUENCECHILD             ParserErrors = 1750
	SCHEMAPUNKNOWNSIMPLECONTENTCHILD        ParserErrors = 1751
	SCHEMAPUNKNOWNSIMPLETYPECHILD           ParserErrors = 1752
	SCHEMAPUNKNOWNTYPE                      ParserErrors = 1753
	SCHEMAPUNKNOWNUNIONCHILD                ParserErrors = 1754
	SCHEMAPELEMDEFAULTFIXED                 ParserErrors = 1755
	SCHEMAPREGEXPINVALID                    ParserErrors = 1756
	SCHEMAPFAILEDLOAD                       ParserErrors = 1757
	SCHEMAPNOTHINGTOPARSE                   ParserErrors = 1758
	SCHEMAPNOROOT                           ParserErrors = 1759
	SCHEMAPREDEFINEDGROUP                   ParserErrors = 1760
	SCHEMAPREDEFINEDTYPE                    ParserErrors = 1761
	SCHEMAPREDEFINEDELEMENT                 ParserErrors = 1762
	SCHEMAPREDEFINEDATTRGROUP               ParserErrors = 1763
	SCHEMAPREDEFINEDATTR                    ParserErrors = 1764
	SCHEMAPREDEFINEDNOTATION                ParserErrors = 1765
	SCHEMAPFAILEDPARSE                      ParserErrors = 1766
	SCHEMAPUNKNOWNPREFIX                    ParserErrors = 1767
	SCHEMAPDEFANDPREFIX                     ParserErrors = 1768
	SCHEMAPUNKNOWNINCLUDECHILD              ParserErrors = 1769
	SCHEMAPINCLUDESCHEMANOTURI              ParserErrors = 1770
	SCHEMAPINCLUDESCHEMANOURI               ParserErrors = 1771
	SCHEMAPNOTSCHEMA                        ParserErrors = 1772
	SCHEMAPUNKNOWNMEMBERTYPE                ParserErrors = 1773
	SCHEMAPINVALIDATTRUSE                   ParserErrors = 1774
	SCHEMAPRECURSIVE                        ParserErrors = 1775
	SCHEMAPSUPERNUMEROUSLISTITEMTYPE        ParserErrors = 1776
	SCHEMAPINVALIDATTRCOMBINATION           ParserErrors = 1777
	SCHEMAPINVALIDATTRINLINECOMBINATION     ParserErrors = 1778
	SCHEMAPMISSINGSIMPLETYPECHILD           ParserErrors = 1779
	SCHEMAPINVALIDATTRNAME                  ParserErrors = 1780
	SCHEMAPREFANDCONTENT                    ParserErrors = 1781
	SCHEMAPCTPROPSCORRECT1                  ParserErrors = 1782
	SCHEMAPCTPROPSCORRECT2                  ParserErrors = 1783
	SCHEMAPCTPROPSCORRECT3                  ParserErrors = 1784
	SCHEMAPCTPROPSCORRECT4                  ParserErrors = 1785
	SCHEMAPCTPROPSCORRECT5                  ParserErrors = 1786
	SCHEMAPDERIVATIONOKRESTRICTION1         ParserErrors = 1787
	SCHEMAPDERIVATIONOKRESTRICTION211       ParserErrors = 1788
	SCHEMAPDERIVATIONOKRESTRICTION212       ParserErrors = 1789
	SCHEMAPDERIVATIONOKRESTRICTION22        ParserErrors = 1790
	SCHEMAPDERIVATIONOKRESTRICTION3         ParserErrors = 1791
	SCHEMAPWILDCARDINVALIDNSMEMBER          ParserErrors = 1792
	SCHEMAPINTERSECTIONNOTEXPRESSIBLE       ParserErrors = 1793
	SCHEMAPUNIONNOTEXPRESSIBLE              ParserErrors = 1794
	SCHEMAPSRCIMPORT31                      ParserErrors = 1795
	SCHEMAPSRCIMPORT32                      ParserErrors = 1796
	SCHEMAPDERIVATIONOKRESTRICTION41        ParserErrors = 1797
	SCHEMAPDERIVATIONOKRESTRICTION42        ParserErrors = 1798
	SCHEMAPDERIVATIONOKRESTRICTION43        ParserErrors = 1799
	SCHEMAPCOSCTEXTENDS13                   ParserErrors = 1800
	SCHEMAVNOROOT                           ParserErrors = 1801
	SCHEMAVUNDECLAREDELEM                   ParserErrors = 1802
	SCHEMAVNOTTOPLEVEL                      ParserErrors = 1803
	SCHEMAVMISSING                          ParserErrors = 1804
	SCHEMAVWRONGELEM                        ParserErrors = 1805
	SCHEMAVNOTYPE                           ParserErrors = 1806
	SCHEMAVNOROLLBACK                       ParserErrors = 1807
	SCHEMAVISABSTRACT                       ParserErrors = 1808
	SCHEMAVNOTEMPTY                         ParserErrors = 1809
	SCHEMAVELEMCONT                         ParserErrors = 1810
	SCHEMAVHAVEDEFAULT                      ParserErrors = 1811
	SCHEMAVNOTNILLABLE                      ParserErrors = 1812
	SCHEMAVEXTRACONTENT                     ParserErrors = 1813
	SCHEMAVINVALIDATTR                      ParserErrors = 1814
	SCHEMAVINVALIDELEM                      ParserErrors = 1815
	SCHEMAVNOTDETERMINIST                   ParserErrors = 1816
	SCHEMAVCONSTRUCT                        ParserErrors = 1817
	SCHEMAVINTERNAL                         ParserErrors = 1818
	SCHEMAVNOTSIMPLE                        ParserErrors = 1819
	SCHEMAVATTRUNKNOWN                      ParserErrors = 1820
	SCHEMAVATTRINVALID                      ParserErrors = 1821
	SCHEMAVVALUE                            ParserErrors = 1822
	SCHEMAVFACET                            ParserErrors = 1823
	SCHEMAVCVCDATATYPEVALID121              ParserErrors = 1824
	SCHEMAVCVCDATATYPEVALID122              ParserErrors = 1825
	SCHEMAVCVCDATATYPEVALID123              ParserErrors = 1826
	SCHEMAVCVCTYPE311                       ParserErrors = 1827
	SCHEMAVCVCTYPE312                       ParserErrors = 1828
	SCHEMAVCVCFACETVALID                    ParserErrors = 1829
	SCHEMAVCVCLENGTHVALID                   ParserErrors = 1830
	SCHEMAVCVCMINLENGTHVALID                ParserErrors = 1831
	SCHEMAVCVCMAXLENGTHVALID                ParserErrors = 1832
	SCHEMAVCVCMININCLUSIVEVALID             ParserErrors = 1833
	SCHEMAVCVCMAXINCLUSIVEVALID             ParserErrors = 1834
	SCHEMAVCVCMINEXCLUSIVEVALID             ParserErrors = 1835
	SCHEMAVCVCMAXEXCLUSIVEVALID             ParserErrors = 1836
	SCHEMAVCVCTOTALDIGITSVALID              ParserErrors = 1837
	SCHEMAVCVCFRACTIONDIGITSVALID           ParserErrors = 1838
	SCHEMAVCVCPATTERNVALID                  ParserErrors = 1839
	SCHEMAVCVCENUMERATIONVALID              ParserErrors = 1840
	SCHEMAVCVCCOMPLEXTYPE21                 ParserErrors = 1841
	SCHEMAVCVCCOMPLEXTYPE22                 ParserErrors = 1842
	SCHEMAVCVCCOMPLEXTYPE23                 ParserErrors = 1843
	SCHEMAVCVCCOMPLEXTYPE24                 ParserErrors = 1844
	SCHEMAVCVCELT1                          ParserErrors = 1845
	SCHEMAVCVCELT2                          ParserErrors = 1846
	SCHEMAVCVCELT31                         ParserErrors = 1847
	SCHEMAVCVCELT321                        ParserErrors = 1848
	SCHEMAVCVCELT322                        ParserErrors = 1849
	SCHEMAVCVCELT41                         ParserErrors = 1850
	SCHEMAVCVCELT42                         ParserErrors = 1851
	SCHEMAVCVCELT43                         ParserErrors = 1852
	SCHEMAVCVCELT511                        ParserErrors = 1853
	SCHEMAVCVCELT512                        ParserErrors = 1854
	SCHEMAVCVCELT521                        ParserErrors = 1855
	SCHEMAVCVCELT5221                       ParserErrors = 1856
	SCHEMAVCVCELT52221                      ParserErrors = 1857
	SCHEMAVCVCELT52222                      ParserErrors = 1858
	SCHEMAVCVCELT6                          ParserErrors = 1859
	SCHEMAVCVCELT7                          ParserErrors = 1860
	SCHEMAVCVCATTRIBUTE1                    ParserErrors = 1861
	SCHEMAVCVCATTRIBUTE2                    ParserErrors = 1862
	SCHEMAVCVCATTRIBUTE3                    ParserErrors = 1863
	SCHEMAVCVCATTRIBUTE4                    ParserErrors = 1864
	SCHEMAVCVCCOMPLEXTYPE31                 ParserErrors = 1865
	SCHEMAVCVCCOMPLEXTYPE321                ParserErrors = 1866
	SCHEMAVCVCCOMPLEXTYPE322                ParserErrors = 1867
	SCHEMAVCVCCOMPLEXTYPE4                  ParserErrors = 1868
	SCHEMAVCVCCOMPLEXTYPE51                 ParserErrors = 1869
	SCHEMAVCVCCOMPLEXTYPE52                 ParserErrors = 1870
	SCHEMAVELEMENTCONTENT                   ParserErrors = 1871
	SCHEMAVDOCUMENTELEMENTMISSING           ParserErrors = 1872
	SCHEMAVCVCCOMPLEXTYPE1                  ParserErrors = 1873
	SCHEMAVCVCAU                            ParserErrors = 1874
	SCHEMAVCVCTYPE1                         ParserErrors = 1875
	SCHEMAVCVCTYPE2                         ParserErrors = 1876
	SCHEMAVCVCIDC                           ParserErrors = 1877
	SCHEMAVCVCWILDCARD                      ParserErrors = 1878
	SCHEMAVMISC                             ParserErrors = 1879
	XPTRUNKNOWNSCHEME                       ParserErrors = 1900
	XPTRCHILDSEQSTART                       ParserErrors = 1901
	XPTREVALFAILED                          ParserErrors = 1902
	XPTREXTRAOBJECTS                        ParserErrors = 1903
	C14NCREATECTXT                          ParserErrors = 1950
	C14NREQUIRESUTF8                        ParserErrors = 1951
	C14NCREATESTACK                         ParserErrors = 1952
	C14NINVALIDNODE                         ParserErrors = 1953
	C14NUNKNOWNODE                          ParserErrors = 1954
	C14NRELATIVENAMESPACE                   ParserErrors = 1955
	FTPPASVANSWER                           ParserErrors = 2000
	FTPEPSVANSWER                           ParserErrors = 2001
	FTPACCNT                                ParserErrors = 2002
	FTPURLSYNTAX                            ParserErrors = 2003
	HTTPURLSYNTAX                           ParserErrors = 2020
	HTTPUSEIP                               ParserErrors = 2021
	HTTPUNKNOWNHOST                         ParserErrors = 2022
	SCHEMAPSRCSIMPLETYPE1                   ParserErrors = 3000
	SCHEMAPSRCSIMPLETYPE2                   ParserErrors = 3001
	SCHEMAPSRCSIMPLETYPE3                   ParserErrors = 3002
	SCHEMAPSRCSIMPLETYPE4                   ParserErrors = 3003
	SCHEMAPSRCRESOLVE                       ParserErrors = 3004
	SCHEMAPSRCRESTRICTIONBASEORSIMPLETYPE   ParserErrors = 3005
	SCHEMAPSRCLISTITEMTYPEORSIMPLETYPE      ParserErrors = 3006
	SCHEMAPSRCUNIONMEMBERTYPESORSIMPLETYPES ParserErrors = 3007
	SCHEMAPSTPROPSCORRECT1                  ParserErrors = 3008
	SCHEMAPSTPROPSCORRECT2                  ParserErrors = 3009
	SCHEMAPSTPROPSCORRECT3                  ParserErrors = 3010
	SCHEMAPCOSSTRESTRICTS11                 ParserErrors = 3011
	SCHEMAPCOSSTRESTRICTS12                 ParserErrors = 3012
	SCHEMAPCOSSTRESTRICTS131                ParserErrors = 3013
	SCHEMAPCOSSTRESTRICTS132                ParserErrors = 3014
	SCHEMAPCOSSTRESTRICTS21                 ParserErrors = 3015
	SCHEMAPCOSSTRESTRICTS2311               ParserErrors = 3016
	SCHEMAPCOSSTRESTRICTS2312               ParserErrors = 3017
	SCHEMAPCOSSTRESTRICTS2321               ParserErrors = 3018
	SCHEMAPCOSSTRESTRICTS2322               ParserErrors = 3019
	SCHEMAPCOSSTRESTRICTS2323               ParserErrors = 3020
	SCHEMAPCOSSTRESTRICTS2324               ParserErrors = 3021
	SCHEMAPCOSSTRESTRICTS2325               ParserErrors = 3022
	SCHEMAPCOSSTRESTRICTS31                 ParserErrors = 3023
	SCHEMAPCOSSTRESTRICTS331                ParserErrors = 3024
	SCHEMAPCOSSTRESTRICTS3312               ParserErrors = 3025
	SCHEMAPCOSSTRESTRICTS3322               ParserErrors = 3026
	SCHEMAPCOSSTRESTRICTS3321               ParserErrors = 3027
	SCHEMAPCOSSTRESTRICTS3323               ParserErrors = 3028
	SCHEMAPCOSSTRESTRICTS3324               ParserErrors = 3029
	SCHEMAPCOSSTRESTRICTS3325               ParserErrors = 3030
	SCHEMAPCOSSTDERIVEDOK21                 ParserErrors = 3031
	SCHEMAPCOSSTDERIVEDOK22                 ParserErrors = 3032
	SCHEMAPS4SELEMNOTALLOWED                ParserErrors = 3033
	SCHEMAPS4SELEMMISSING                   ParserErrors = 3034
	SCHEMAPS4SATTRNOTALLOWED                ParserErrors = 3035
	SCHEMAPS4SATTRMISSING                   ParserErrors = 3036
	SCHEMAPS4SATTRINVALIDVALUE              ParserErrors = 3037
	SCHEMAPSRCELEMENT1                      ParserErrors = 3038
	SCHEMAPSRCELEMENT21                     ParserErrors = 3039
	SCHEMAPSRCELEMENT22                     ParserErrors = 3040
	SCHEMAPSRCELEMENT3                      ParserErrors = 3041
	SCHEMAPPPROPSCORRECT1                   ParserErrors = 3042
	SCHEMAPPPROPSCORRECT21                  ParserErrors = 3043
	SCHEMAPPPROPSCORRECT22                  ParserErrors = 3044
	SCHEMAPEPROPSCORRECT2                   ParserErrors = 3045
	SCHEMAPEPROPSCORRECT3                   ParserErrors = 3046
	SCHEMAPEPROPSCORRECT4                   ParserErrors = 3047
	SCHEMAPEPROPSCORRECT5                   ParserErrors = 3048
	SCHEMAPEPROPSCORRECT6                   ParserErrors = 3049
	SCHEMAPSRCINCLUDE                       ParserErrors = 3050
	SCHEMAPSRCATTRIBUTE1                    ParserErrors = 3051
	SCHEMAPSRCATTRIBUTE2                    ParserErrors = 3052
	SCHEMAPSRCATTRIBUTE31                   ParserErrors = 3053
	SCHEMAPSRCATTRIBUTE32                   ParserErrors = 3054
	SCHEMAPSRCATTRIBUTE4                    ParserErrors = 3055
	SCHEMAPNOXMLNS                          ParserErrors = 3056
	SCHEMAPNOXSI                            ParserErrors = 3057
	SCHEMAPCOSVALIDDEFAULT1                 ParserErrors = 3058
	SCHEMAPCOSVALIDDEFAULT21                ParserErrors = 3059
	SCHEMAPCOSVALIDDEFAULT221               ParserErrors = 3060
	SCHEMAPCOSVALIDDEFAULT222               ParserErrors = 3061
	SCHEMAPCVCSIMPLETYPE                    ParserErrors = 3062
	SCHEMAPCOSCTEXTENDS11                   ParserErrors = 3063
	SCHEMAPSRCIMPORT11                      ParserErrors = 3064
	SCHEMAPSRCIMPORT12                      ParserErrors = 3065
	SCHEMAPSRCIMPORT2                       ParserErrors = 3066
	SCHEMAPSRCIMPORT21                      ParserErrors = 3067
	SCHEMAPSRCIMPORT22                      ParserErrors = 3068
	SCHEMAPINTERNAL                         ParserErrors = 3069
	SCHEMAPNOTDETERMINISTIC                 ParserErrors = 3070
	SCHEMAPSRCATTRIBUTEGROUP1               ParserErrors = 3071
	SCHEMAPSRCATTRIBUTEGROUP2               ParserErrors = 3072
	SCHEMAPSRCATTRIBUTEGROUP3               ParserErrors = 3073
	SCHEMAPMGPROPSCORRECT1                  ParserErrors = 3074
	SCHEMAPMGPROPSCORRECT2                  ParserErrors = 3075
	SCHEMAPSRCCT1                           ParserErrors = 3076
	SCHEMAPDERIVATIONOKRESTRICTION213       ParserErrors = 3077
	SCHEMAPAUPROPSCORRECT2                  ParserErrors = 3078
	SCHEMAPAPROPSCORRECT2                   ParserErrors = 3079
	SCHEMAPCPROPSCORRECT                    ParserErrors = 3080
	SCHEMAPSRCREDEFINE                      ParserErrors = 3081
	SCHEMAPSRCIMPORT                        ParserErrors = 3082
	SCHEMAPWARNSKIPSCHEMA                   ParserErrors = 3083
	SCHEMAPWARNUNLOCATEDSCHEMA              ParserErrors = 3084
	SCHEMAPWARNATTRREDECLPROH               ParserErrors = 3085
	SCHEMAPWARNATTRPOINTLESSPROH            ParserErrors = 3086
	SCHEMAPAGPROPSCORRECT                   ParserErrors = 3087
	SCHEMAPCOSCTEXTENDS12                   ParserErrors = 3088
	SCHEMAPAUPROPSCORRECT                   ParserErrors = 3089
	SCHEMAPAPROPSCORRECT3                   ParserErrors = 3090
	SCHEMAPCOSALLLIMITED                    ParserErrors = 3091
	SCHEMATRONVASSERT                       ParserErrors = 4000
	SCHEMATRONVREPORT                       ParserErrors = 4001
	MODULEOPEN                              ParserErrors = 4900
	MODULECLOSE                             ParserErrors = 4901
	CHECKFOUNDELEMENT                       ParserErrors = 5000
	CHECKFOUNDATTRIBUTE                     ParserErrors = 5001
	CHECKFOUNDTEXT                          ParserErrors = 5002
	CHECKFOUNDCDATA                         ParserErrors = 5003
	CHECKFOUNDENTITYREF                     ParserErrors = 5004
	CHECKFOUNDENTITY                        ParserErrors = 5005
	CHECKFOUNDPI                            ParserErrors = 5006
	CHECKFOUNDCOMMENT                       ParserErrors = 5007
	CHECKFOUNDDOCTYPE                       ParserErrors = 5008
	CHECKFOUNDFRAGMENT                      ParserErrors = 5009
	CHECKFOUNDNOTATION                      ParserErrors = 5010
	CHECKUNKNOWNNODE                        ParserErrors = 5011
	CHECKENTITYTYPE                         ParserErrors = 5012
	CHECKNOPARENT                           ParserErrors = 5013
	CHECKNODOC                              ParserErrors = 5014
	CHECKNONAME                             ParserErrors = 5015
	CHECKNOELEM                             ParserErrors = 5016
	CHECKWRONGDOC                           ParserErrors = 5017
	CHECKNOPREV                             ParserErrors = 5018
	CHECKWRONGPREV                          ParserErrors = 5019
	CHECKNONEXT                             ParserErrors = 5020
	CHECKWRONGNEXT                          ParserErrors = 5021
	CHECKNOTDTD                             ParserErrors = 5022
	CHECKNOTATTR                            ParserErrors = 5023
	CHECKNOTATTRDECL                        ParserErrors = 5024
	CHECKNOTELEMDECL                        ParserErrors = 5025
	CHECKNOTENTITYDECL                      ParserErrors = 5026
	CHECKNOTNSDECL                          ParserErrors = 5027
	CHECKNOHREF                             ParserErrors = 5028
	CHECKWRONGPARENT                        ParserErrors = 5029
	CHECKNSSCOPE                            ParserErrors = 5030
	CHECKNSANCESTOR                         ParserErrors = 5031
	CHECKNOTUTF8                            ParserErrors = 5032
	CHECKNODICT                             ParserErrors = 5033
	CHECKNOTNCNAME                          ParserErrors = 5034
	CHECKOUTSIDEDICT                        ParserErrors = 5035
	CHECKWRONGNAME                          ParserErrors = 5036
	CHECKNAMENOTNULL                        ParserErrors = 5037
	I18NNONAME                              ParserErrors = 6000
	I18NNOHANDLER                           ParserErrors = 6001
	I18NEXCESSHANDLER                       ParserErrors = 6002
	I18NCONVFAILED                          ParserErrors = 6003
	I18NNOOUTPUT                            ParserErrors = 6004
	BUFOVERFLOW                             ParserErrors = 7000
)

// llgo:type C
type GenericErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type StructuredErrorFunc func(unsafe.Pointer, *Error)

//go:linkname X__xmlLastError C.__xmlLastError
func X__xmlLastError() *Error

//go:linkname X__xmlGenericError C.__xmlGenericError
func X__xmlGenericError() GenericErrorFunc

//go:linkname X__xmlGenericErrorContext C.__xmlGenericErrorContext
func X__xmlGenericErrorContext() *unsafe.Pointer

//go:linkname X__xmlStructuredError C.__xmlStructuredError
func X__xmlStructuredError() StructuredErrorFunc

//go:linkname X__xmlStructuredErrorContext C.__xmlStructuredErrorContext
func X__xmlStructuredErrorContext() *unsafe.Pointer

/*
 * Use the following function to reset the two global variables
 * xmlGenericError and xmlGenericErrorContext.
 */
//go:linkname SetGenericErrorFunc C.xmlSetGenericErrorFunc
func SetGenericErrorFunc(ctx unsafe.Pointer, handler GenericErrorFunc)

//go:linkname ThrDefSetGenericErrorFunc C.xmlThrDefSetGenericErrorFunc
func ThrDefSetGenericErrorFunc(ctx unsafe.Pointer, handler GenericErrorFunc)

//go:linkname InitGenericErrorDefaultFunc C.initGenericErrorDefaultFunc
func InitGenericErrorDefaultFunc(handler GenericErrorFunc)

//go:linkname SetStructuredErrorFunc C.xmlSetStructuredErrorFunc
func SetStructuredErrorFunc(ctx unsafe.Pointer, handler StructuredErrorFunc)

//go:linkname ThrDefSetStructuredErrorFunc C.xmlThrDefSetStructuredErrorFunc
func ThrDefSetStructuredErrorFunc(ctx unsafe.Pointer, handler StructuredErrorFunc)

/*
 * Default message routines used by SAX and Valid context for error
 * and warning reporting.
 */
//go:linkname ParserError C.xmlParserError
func ParserError(ctx unsafe.Pointer, msg *int8, __llgo_va_list ...interface{})

//go:linkname ParserWarning C.xmlParserWarning
func ParserWarning(ctx unsafe.Pointer, msg *int8, __llgo_va_list ...interface{})

//go:linkname ParserValidityError C.xmlParserValidityError
func ParserValidityError(ctx unsafe.Pointer, msg *int8, __llgo_va_list ...interface{})

//go:linkname ParserValidityWarning C.xmlParserValidityWarning
func ParserValidityWarning(ctx unsafe.Pointer, msg *int8, __llgo_va_list ...interface{})

/** DOC_ENABLE */
// llgo:link (*X_xmlParserInput).ParserPrintFileInfo C.xmlParserPrintFileInfo
func (recv_ *X_xmlParserInput) ParserPrintFileInfo() {
}

// llgo:link (*X_xmlParserInput).ParserPrintFileContext C.xmlParserPrintFileContext
func (recv_ *X_xmlParserInput) ParserPrintFileContext() {
}

// llgo:link (*Error).FormatError C.xmlFormatError
func (recv_ *Error) FormatError(channel GenericErrorFunc, data unsafe.Pointer) {
}

/*
 * Extended error information routines
 */
//go:linkname GetLastError C.xmlGetLastError
func GetLastError() *Error

//go:linkname ResetLastError C.xmlResetLastError
func ResetLastError()

//go:linkname CtxtGetLastError C.xmlCtxtGetLastError
func CtxtGetLastError(ctx unsafe.Pointer) *Error

//go:linkname CtxtResetLastError C.xmlCtxtResetLastError
func CtxtResetLastError(ctx unsafe.Pointer)

//go:linkname ResetError C.xmlResetError
func ResetError(err ErrorPtr)

// llgo:link (*Error).CopyError C.xmlCopyError
func (recv_ *Error) CopyError(to ErrorPtr) c.Int {
	return 0
}
