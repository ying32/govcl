unit MSTSCLib_TLB;

// ************************************************************************ //
// WARNING                                                                    
// -------                                                                    
// The types declared in this file were generated from data read from a       
// Type Library. If this type library is explicitly or indirectly (via        
// another type library referring to this type library) re-imported, or the   
// 'Refresh' command of the Type Library Editor activated while editing the   
// Type Library, the contents of this file will be regenerated and all        
// manual modifications will be lost.                                         
// ************************************************************************ //

// $Rev: 52393 $
// File generated on 2019/6/21 8:32:02 from Type Library described below.

// ************************************************************************  //
// Type Lib: C:\Windows\SysWOW64\mstscax.dll (1)
// LIBID: {8C11EFA1-92C3-11D1-BC1E-00C04FA31489}
// LCID: 0
// Helpfile: 
// HelpString: Microsoft Terminal Services Control Type Library
// DepndLst: 
//   (1) v2.0 stdole, (C:\Windows\SysWOW64\stdole2.tlb)
// SYS_KIND: SYS_WIN32
// Errors:
//   Hint: Member 'Property' of 'IMsRdpExtendedSettings' changed to 'Property_'
//   Error creating palette bitmap of (TMsTscAxNotSafeForScripting) : Registry entry for control bitmap is invalid
//   Error creating palette bitmap of (TMsRdpClientNotSafeForScripting) : Registry entry for control bitmap is invalid
//   Error creating palette bitmap of (TMsRdpClient2NotSafeForScripting) : Registry entry for control bitmap is invalid
//   Error creating palette bitmap of (TMsRdpClient9NotSafeForScripting) : Registry key CLSID\{8B918B82-7985-4C24-89DF-C33AD2BBFBCD}\ToolboxBitmap32 not found
//   Error creating palette bitmap of (TMsRdpClient9) : Registry key CLSID\{301B94BA-5D25-4A12-BFFE-3B6E7A616585}\ToolboxBitmap32 not found
// ************************************************************************ //
{$TYPEDADDRESS OFF} // Unit must be compiled without type-checked pointers. 
{$WARN SYMBOL_PLATFORM OFF}
{$WRITEABLECONST ON}
{$VARPROPSETTER ON}
{$ALIGN 4}

interface

uses Winapi.Windows, System.Classes, System.Variants, System.Win.StdVCL, Vcl.Graphics, Vcl.OleCtrls, Vcl.OleServer, Winapi.ActiveX;
  


// *********************************************************************//
// GUIDS declared in the TypeLibrary. Following prefixes are used:        
//   Type Libraries     : LIBID_xxxx                                      
//   CoClasses          : CLASS_xxxx                                      
//   DISPInterfaces     : DIID_xxxx                                       
//   Non-DISP interfaces: IID_xxxx                                        
// *********************************************************************//
const
  // TypeLibrary Major and minor versions
  MSTSCLibMajorVersion = 1;
  MSTSCLibMinorVersion = 0;

  LIBID_MSTSCLib: TGUID = '{8C11EFA1-92C3-11D1-BC1E-00C04FA31489}';

  DIID_IMsTscAxEvents: TGUID = '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
  IID_IMsTscAx_Redist: TGUID = '{327BB5CD-834E-4400-AEF2-B30E15E5D682}';
  IID_IMsTscAx: TGUID = '{8C11EFAE-92C3-11D1-BC1E-00C04FA31489}';
  IID_IMsRdpClient: TGUID = '{92B4A539-7115-4B7C-A5A9-E5D9EFC2780A}';
  IID_IMsTscNonScriptable: TGUID = '{C1E6743A-41C1-4A74-832A-0DD06C1C7A0E}';
  IID_IMsTscSecuredSettings: TGUID = '{C9D65442-A0F9-45B2-8F73-D61D2DB8CBB6}';
  IID_IMsTscAdvancedSettings: TGUID = '{809945CC-4B3B-4A92-A6B0-DBF9B5F2EF2D}';
  IID_IMsTscDebug: TGUID = '{209D0EB9-6254-47B1-9033-A98DAE55BB27}';
  IID_IMsRdpClientAdvancedSettings: TGUID = '{3C65B4AB-12B3-465B-ACD4-B8DAD3BFF9E2}';
  IID_IMsRdpClientSecuredSettings: TGUID = '{605BEFCF-39C1-45CC-A811-068FB7BE346D}';
  IID_IMsRdpClientNonScriptable: TGUID = '{2F079C4C-87B2-4AFD-97AB-20CDB43038AE}';
  CLASS_MsTscAxNotSafeForScripting: TGUID = '{A41A4187-5A86-4E26-B40A-856F9035D9CB}';
  CLASS_MsTscAx: TGUID = '{1FB464C8-09BB-4017-A2F5-EB742F04392F}';
  CLASS_MsRdpClientNotSafeForScripting: TGUID = '{7CACBD7B-0D99-468F-AC33-22E495C0AFE5}';
  CLASS_MsRdpClient: TGUID = '{791FA017-2DE3-492E-ACC5-53C67A2B94D0}';
  IID_IMsRdpClient2: TGUID = '{E7E17DC4-3B71-4BA7-A8E6-281FFADCA28F}';
  CLASS_MsRdpClient2NotSafeForScripting: TGUID = '{3523C2FB-4031-44E4-9A3B-F1E94986EE7F}';
  IID_IMsRdpClientAdvancedSettings2: TGUID = '{9AC42117-2B76-4320-AA44-0E616AB8437B}';
  CLASS_MsRdpClient2: TGUID = '{9059F30F-4EB1-4BD2-9FDC-36F43A218F4A}';
  CLASS_MsRdpClient2a: TGUID = '{971127BB-259F-48C2-BD75-5F97A3331551}';
  IID_IMsRdpClient3: TGUID = '{91B7CBC5-A72E-4FA0-9300-D647D7E897FF}';
  CLASS_MsRdpClient3NotSafeForScripting: TGUID = '{ACE575FD-1FCF-4074-9401-EBAB990FA9DE}';
  IID_IMsRdpClientAdvancedSettings3: TGUID = '{19CD856B-C542-4C53-ACEE-F127E3BE1A59}';
  CLASS_MsRdpClient3: TGUID = '{7584C670-2274-4EFB-B00B-D6AABA6D3850}';
  CLASS_MsRdpClient3a: TGUID = '{6A6F4B83-45C5-4CA9-BDD9-0D81C12295E4}';
  IID_IMsRdpClient4: TGUID = '{095E0738-D97D-488B-B9F6-DD0E8D66C0DE}';
  IID_IMsRdpClientNonScriptable2: TGUID = '{17A5E535-4072-4FA4-AF32-C8D0D47345E9}';
  IID_IMsRdpClientAdvancedSettings4: TGUID = '{FBA7F64E-7345-4405-AE50-FA4A763DC0DE}';
  CLASS_MsRdpClient4NotSafeForScripting: TGUID = '{6AE29350-321B-42BE-BBE5-12FB5270C0DE}';
  CLASS_MsRdpClient4: TGUID = '{4EDCB26C-D24C-4E72-AF07-B576699AC0DE}';
  CLASS_MsRdpClient4a: TGUID = '{54CE37E0-9834-41AE-9896-4DAB69DC022B}';
  IID_IMsRdpClient5: TGUID = '{4EB5335B-6429-477D-B922-E06A28ECD8BF}';
  IID_IMsRdpClientNonScriptable3: TGUID = '{B3378D90-0728-45C7-8ED7-B6159FB92219}';
  IID_IMsRdpClientTransportSettings: TGUID = '{720298C0-A099-46F5-9F82-96921BAE4701}';
  IID_IMsRdpClientAdvancedSettings5: TGUID = '{FBA7F64E-6783-4405-DA45-FA4A763DABD0}';
  IID_ITSRemoteProgram: TGUID = '{FDD029F9-467A-4C49-8529-64B521DBD1B4}';
  IID_IMsRdpClientShell: TGUID = '{D012AE6D-C19A-4BFE-B367-201F8911F134}';
  CLASS_MsRdpClient5NotSafeForScripting: TGUID = '{4EB2F086-C818-447E-B32C-C51CE2B30D31}';
  IID_IMsRdpDeviceCollection: TGUID = '{56540617-D281-488C-8738-6A8FDF64A118}';
  IID_IMsRdpDevice: TGUID = '{60C3B9C8-9E92-4F5E-A3E7-604A912093EA}';
  IID_IMsRdpDriveCollection: TGUID = '{7FF17599-DA2C-4677-AD35-F60C04FE1585}';
  IID_IMsRdpDrive: TGUID = '{D28B5458-F694-47A8-8E61-40356A767E46}';
  CLASS_MsRdpClient5: TGUID = '{4EB89FF4-7F78-4A0F-8B8D-2BF02E94E4B2}';
  IID_IMsRdpClient6: TGUID = '{D43B7D80-8517-4B6D-9EAC-96AD6800D7F2}';
  IID_IMsRdpClientNonScriptable4: TGUID = '{F50FA8AA-1C7D-4F59-B15C-A90CACAE1FCB}';
  IID_IMsRdpClientAdvancedSettings6: TGUID = '{222C4B5D-45D9-4DF0-A7C6-60CF9089D285}';
  IID_IMsRdpClientTransportSettings2: TGUID = '{67341688-D606-4C73-A5D2-2E0489009319}';
  CLASS_MsRdpClient6NotSafeForScripting: TGUID = '{D2EA46A7-C2BF-426B-AF24-E19C44456399}';
  CLASS_MsRdpClient6: TGUID = '{7390F3D8-0439-4C05-91E3-CF5CB290C3D0}';
  IID_IMsRdpClient7: TGUID = '{B2A5B5CE-3461-444A-91D4-ADD26D070638}';
  IID_IMsRdpClientNonScriptable5: TGUID = '{4F6996D5-D7B1-412C-B0FF-063718566907}';
  IID_IMsRdpClientAdvancedSettings7: TGUID = '{26036036-4010-4578-8091-0DB9A1EDF9C3}';
  IID_IMsRdpClientTransportSettings3: TGUID = '{3D5B21AC-748D-41DE-8F30-E15169586BD4}';
  IID_IMsRdpClientSecuredSettings2: TGUID = '{25F2CE20-8B1D-4971-A7CD-549DAE201FC0}';
  IID_ITSRemoteProgram2: TGUID = '{92C38A7D-241A-418C-9936-099872C9AF20}';
  IID_IMsRdpPreferredRedirectionInfo: TGUID = '{FDD029F9-9574-4DEF-8529-64B521CCCAA4}';
  IID_IMsRdpExtendedSettings: TGUID = '{302D8188-0052-4807-806A-362B628F9AC5}';
  CLASS_MsRdpClient7NotSafeForScripting: TGUID = '{54D38BF7-B1EF-4479-9674-1BD6EA465258}';
  CLASS_MsRdpClient7: TGUID = '{A9D7038D-B5ED-472E-9C47-94BEA90A5910}';
  IID_IMsRdpClient8: TGUID = '{4247E044-9271-43A9-BC49-E2AD9E855D62}';
  CLASS_MsRdpClient8NotSafeForScripting: TGUID = '{A3BC03A0-041D-42E3-AD22-882B7865C9C5}';
  IID_IMsRdpClientAdvancedSettings8: TGUID = '{89ACB528-2557-4D16-8625-226A30E97E9A}';
  CLASS_MsRdpClient8: TGUID = '{5F681803-2900-4C43-A1CC-CF405404A676}';
  IID_IMsRdpClient9: TGUID = '{28904001-04B6-436C-A55B-0AF1A0883DC9}';
  CLASS_MsRdpClient9NotSafeForScripting: TGUID = '{8B918B82-7985-4C24-89DF-C33AD2BBFBCD}';
  IID_IMsRdpClientTransportSettings4: TGUID = '{011C3236-4D81-4515-9143-067AB630D299}';
  CLASS_MsRdpClient9: TGUID = '{301B94BA-5D25-4A12-BFFE-3B6E7A616585}';
  DIID_IRemoteDesktopClientEvents: TGUID = '{079863B7-6D47-4105-8BFE-0CDCB360E67D}';
  IID_IRemoteDesktopClient: TGUID = '{57D25668-625A-4905-BE4E-304CAA13F89C}';
  CLASS_RemoteDesktopClient: TGUID = '{EAB16C5D-EED1-4E95-868B-0FBA1B42C092}';
  IID_IRemoteDesktopClientSettings: TGUID = '{48A0F2A7-2713-431F-BBAC-6F4558E7D64D}';
  IID_IRemoteDesktopClientActions: TGUID = '{7D54BC4E-1028-45D4-8B0A-B9B6BFFBA176}';
  IID_IRemoteDesktopClientTouchPointer: TGUID = '{260EC22D-8CBC-44B5-9E88-2A37F6C93AE9}';

// *********************************************************************//
// Declaration of Enumerations defined in Type Library                    
// *********************************************************************//
// Constants for enum __MIDL___MIDL_itf_mstsax_0000_0037_0001
type
  __MIDL___MIDL_itf_mstsax_0000_0037_0001 = TOleEnum;
const
  autoReconnectContinueAutomatic = $00000000;
  autoReconnectContinueStop = $00000001;
  autoReconnectContinueManual = $00000002;

// Constants for enum __MIDL___MIDL_itf_mstsax_0000_0049_0001
type
  __MIDL___MIDL_itf_mstsax_0000_0049_0001 = TOleEnum;
const
  remoteAppResultOk = $00000000;
  remoteAppResultLocked = $00000001;
  remoteAppResultProtocolError = $00000002;
  remoteAppResultNotInWhitelist = $00000003;
  remoteAppResultNetworkPathDenied = $00000004;
  remoteAppResultFileNotFound = $00000005;
  remoteAppResultFailure = $00000006;
  remoteAppResultHookNotLoaded = $00000007;

// Constants for enum __MIDL___MIDL_itf_mstsax_0000_0049_0002
type
  __MIDL___MIDL_itf_mstsax_0000_0049_0002 = TOleEnum;
const
  remoteAppWindowNone = $00000000;
  remoteAppWindowDisplayed = $00000001;
  remoteAppShellIconDisplayed = $00000002;

// Constants for enum __MIDL___MIDL_itf_mstsax_0000_0000_0001
type
  __MIDL___MIDL_itf_mstsax_0000_0000_0001 = TOleEnum;
const
  exDiscReasonNoInfo = $00000000;
  exDiscReasonAPIInitiatedDisconnect = $00000001;
  exDiscReasonAPIInitiatedLogoff = $00000002;
  exDiscReasonServerIdleTimeout = $00000003;
  exDiscReasonServerLogonTimeout = $00000004;
  exDiscReasonReplacedByOtherConnection = $00000005;
  exDiscReasonOutOfMemory = $00000006;
  exDiscReasonServerDeniedConnection = $00000007;
  exDiscReasonServerDeniedConnectionFips = $00000008;
  exDiscReasonServerInsufficientPrivileges = $00000009;
  exDiscReasonServerFreshCredsRequired = $0000000A;
  exDiscReasonRpcInitiatedDisconnectByUser = $0000000B;
  exDiscReasonLogoffByUser = $0000000C;
  exDiscReasonLicenseInternal = $00000100;
  exDiscReasonLicenseNoLicenseServer = $00000101;
  exDiscReasonLicenseNoLicense = $00000102;
  exDiscReasonLicenseErrClientMsg = $00000103;
  exDiscReasonLicenseHwidDoesntMatchLicense = $00000104;
  exDiscReasonLicenseErrClientLicense = $00000105;
  exDiscReasonLicenseCantFinishProtocol = $00000106;
  exDiscReasonLicenseClientEndedProtocol = $00000107;
  exDiscReasonLicenseErrClientEncryption = $00000108;
  exDiscReasonLicenseCantUpgradeLicense = $00000109;
  exDiscReasonLicenseNoRemoteConnections = $0000010A;
  exDiscReasonLicenseCreatingLicStoreAccDenied = $0000010B;
  exDiscReasonRdpEncInvalidCredentials = $00000300;
  exDiscReasonProtocolRangeStart = $00001000;
  exDiscReasonProtocolRangeEnd = $00007FFF;

// Constants for enum __MIDL_IMsRdpClient_0001
type
  __MIDL_IMsRdpClient_0001 = TOleEnum;
const
  controlCloseCanProceed = $00000000;
  controlCloseWaitForEvents = $00000001;

// Constants for enum __MIDL_IMsRdpClientNonScriptable4_0001
type
  __MIDL_IMsRdpClientNonScriptable4_0001 = TOleEnum;
const
  RedirectionWarningTypeDefault = $00000000;
  RedirectionWarningTypeUnsigned = $00000001;
  RedirectionWarningTypeUnknown = $00000002;
  RedirectionWarningTypeUser = $00000003;
  RedirectionWarningTypeThirdPartySigned = $00000004;
  RedirectionWarningTypeTrusted = $00000005;
  RedirectionWarningTypeMax = $00000005;

// Constants for enum __MIDL___MIDL_itf_mstsax_0000_0000_0004
type
  __MIDL___MIDL_itf_mstsax_0000_0000_0004 = TOleEnum;
const
  RemoteSessionActionCharms = $00000000;
  RemoteSessionActionAppbar = $00000001;
  RemoteSessionActionSnap = $00000002;
  RemoteSessionActionStartScreen = $00000003;
  RemoteSessionActionAppSwitch = $00000004;

// Constants for enum __MIDL___MIDL_itf_mstsax_0000_0000_0003
type
  __MIDL___MIDL_itf_mstsax_0000_0000_0003 = TOleEnum;
const
  FullMode = $00000000;
  ThinClientMode = $00000001;
  SmallCacheMode = $00000002;

// Constants for enum __MIDL_IMsRdpClient8_0001
type
  __MIDL_IMsRdpClient8_0001 = TOleEnum;
const
  controlReconnectStarted = $00000000;
  controlReconnectBlocked = $00000001;

// Constants for enum __MIDL_IRemoteDesktopClientActions_0001
type
  __MIDL_IRemoteDesktopClientActions_0001 = TOleEnum;
const
  RemoteActionCharms = $00000000;
  RemoteActionAppbar = $00000001;
  RemoteActionSnap = $00000002;
  RemoteActionStartScreen = $00000003;
  RemoteActionAppSwitch = $00000004;

// Constants for enum __MIDL_IRemoteDesktopClientActions_0002
type
  __MIDL_IRemoteDesktopClientActions_0002 = TOleEnum;
const
  SnapshotEncodingDataUri = $00000000;

// Constants for enum __MIDL_IRemoteDesktopClientActions_0003
type
  __MIDL_IRemoteDesktopClientActions_0003 = TOleEnum;
const
  SnapshotFormatPng = $00000000;
  SnapshotFormatJpeg = $00000001;
  SnapshotFormatBmp = $00000002;

type

// *********************************************************************//
// Forward declaration of types defined in TypeLibrary                    
// *********************************************************************//
  IMsTscAxEvents = dispinterface;
  IMsTscAx_Redist = interface;
  IMsTscAx_RedistDisp = dispinterface;
  IMsTscAx = interface;
  IMsTscAxDisp = dispinterface;
  IMsRdpClient = interface;
  IMsRdpClientDisp = dispinterface;
  IMsTscNonScriptable = interface;
  IMsTscSecuredSettings = interface;
  IMsTscSecuredSettingsDisp = dispinterface;
  IMsTscAdvancedSettings = interface;
  IMsTscAdvancedSettingsDisp = dispinterface;
  IMsTscDebug = interface;
  IMsTscDebugDisp = dispinterface;
  IMsRdpClientAdvancedSettings = interface;
  IMsRdpClientAdvancedSettingsDisp = dispinterface;
  IMsRdpClientSecuredSettings = interface;
  IMsRdpClientSecuredSettingsDisp = dispinterface;
  IMsRdpClientNonScriptable = interface;
  IMsRdpClient2 = interface;
  IMsRdpClient2Disp = dispinterface;
  IMsRdpClientAdvancedSettings2 = interface;
  IMsRdpClientAdvancedSettings2Disp = dispinterface;
  IMsRdpClient3 = interface;
  IMsRdpClient3Disp = dispinterface;
  IMsRdpClientAdvancedSettings3 = interface;
  IMsRdpClientAdvancedSettings3Disp = dispinterface;
  IMsRdpClient4 = interface;
  IMsRdpClient4Disp = dispinterface;
  IMsRdpClientNonScriptable2 = interface;
  IMsRdpClientAdvancedSettings4 = interface;
  IMsRdpClientAdvancedSettings4Disp = dispinterface;
  IMsRdpClient5 = interface;
  IMsRdpClient5Disp = dispinterface;
  IMsRdpClientNonScriptable3 = interface;
  IMsRdpClientTransportSettings = interface;
  IMsRdpClientTransportSettingsDisp = dispinterface;
  IMsRdpClientAdvancedSettings5 = interface;
  IMsRdpClientAdvancedSettings5Disp = dispinterface;
  ITSRemoteProgram = interface;
  ITSRemoteProgramDisp = dispinterface;
  IMsRdpClientShell = interface;
  IMsRdpClientShellDisp = dispinterface;
  IMsRdpDeviceCollection = interface;
  IMsRdpDevice = interface;
  IMsRdpDriveCollection = interface;
  IMsRdpDrive = interface;
  IMsRdpClient6 = interface;
  IMsRdpClient6Disp = dispinterface;
  IMsRdpClientNonScriptable4 = interface;
  IMsRdpClientAdvancedSettings6 = interface;
  IMsRdpClientAdvancedSettings6Disp = dispinterface;
  IMsRdpClientTransportSettings2 = interface;
  IMsRdpClientTransportSettings2Disp = dispinterface;
  IMsRdpClient7 = interface;
  IMsRdpClient7Disp = dispinterface;
  IMsRdpClientNonScriptable5 = interface;
  IMsRdpClientAdvancedSettings7 = interface;
  IMsRdpClientAdvancedSettings7Disp = dispinterface;
  IMsRdpClientTransportSettings3 = interface;
  IMsRdpClientTransportSettings3Disp = dispinterface;
  IMsRdpClientSecuredSettings2 = interface;
  IMsRdpClientSecuredSettings2Disp = dispinterface;
  ITSRemoteProgram2 = interface;
  ITSRemoteProgram2Disp = dispinterface;
  IMsRdpPreferredRedirectionInfo = interface;
  IMsRdpExtendedSettings = interface;
  IMsRdpClient8 = interface;
  IMsRdpClient8Disp = dispinterface;
  IMsRdpClientAdvancedSettings8 = interface;
  IMsRdpClientAdvancedSettings8Disp = dispinterface;
  IMsRdpClient9 = interface;
  IMsRdpClient9Disp = dispinterface;
  IMsRdpClientTransportSettings4 = interface;
  IMsRdpClientTransportSettings4Disp = dispinterface;
  IRemoteDesktopClientEvents = dispinterface;
  IRemoteDesktopClient = interface;
  IRemoteDesktopClientDisp = dispinterface;
  IRemoteDesktopClientSettings = interface;
  IRemoteDesktopClientSettingsDisp = dispinterface;
  IRemoteDesktopClientActions = interface;
  IRemoteDesktopClientActionsDisp = dispinterface;
  IRemoteDesktopClientTouchPointer = interface;
  IRemoteDesktopClientTouchPointerDisp = dispinterface;

// *********************************************************************//
// Declaration of CoClasses defined in Type Library                       
// (NOTE: Here we map each CoClass to its Default Interface)              
// *********************************************************************//
  MsTscAxNotSafeForScripting = IMsTscAx;
  MsTscAx = IMsTscAx;
  MsRdpClientNotSafeForScripting = IMsRdpClient;
  MsRdpClient = IMsRdpClient;
  MsRdpClient2NotSafeForScripting = IMsRdpClient2;
  MsRdpClient2 = IMsRdpClient2;
  MsRdpClient2a = IMsRdpClient2;
  MsRdpClient3NotSafeForScripting = IMsRdpClient3;
  MsRdpClient3 = IMsRdpClient3;
  MsRdpClient3a = IMsRdpClient3;
  MsRdpClient4NotSafeForScripting = IMsRdpClient4;
  MsRdpClient4 = IMsRdpClient4;
  MsRdpClient4a = IMsRdpClient4;
  MsRdpClient5NotSafeForScripting = IMsRdpClient5;
  MsRdpClient5 = IMsRdpClient5;
  MsRdpClient6NotSafeForScripting = IMsRdpClient6;
  MsRdpClient6 = IMsRdpClient6;
  MsRdpClient7NotSafeForScripting = IMsRdpClient7;
  MsRdpClient7 = IMsRdpClient7;
  MsRdpClient8NotSafeForScripting = IMsRdpClient8;
  MsRdpClient8 = IMsRdpClient8;
  MsRdpClient9NotSafeForScripting = IMsRdpClient9;
  MsRdpClient9 = IMsRdpClient9;
  RemoteDesktopClient = IRemoteDesktopClient;


// *********************************************************************//
// Declaration of structures, unions and aliases.                         
// *********************************************************************//
  wireHWND = ^_RemotableHandle; 
  POleVariant1 = ^OleVariant; {*}
  PWordBool1 = ^WordBool; {*}
  PInteger1 = ^Integer; {*}

  AutoReconnectContinueState = __MIDL___MIDL_itf_mstsax_0000_0037_0001; 
  RemoteProgramResult = __MIDL___MIDL_itf_mstsax_0000_0049_0001; 

  __MIDL_IWinTypes_0009 = record
    case Integer of
      0: (hInproc: Integer);
      1: (hRemote: Integer);
  end;

  _RemotableHandle = record
    fContext: Integer;
    u: __MIDL_IWinTypes_0009;
  end;

  RemoteWindowDisplayedAttribute = __MIDL___MIDL_itf_mstsax_0000_0049_0002; 
  ExtendedDisconnectReasonCode = __MIDL___MIDL_itf_mstsax_0000_0000_0001; 
  ControlCloseStatus = __MIDL_IMsRdpClient_0001; 
  UINT_PTR = LongWord; 
  LONG_PTR = Integer; 
  RedirectionWarningType = __MIDL_IMsRdpClientNonScriptable4_0001; 
  RemoteSessionActionType = __MIDL___MIDL_itf_mstsax_0000_0000_0004; 
  ClientSpec = __MIDL___MIDL_itf_mstsax_0000_0000_0003; 
  ControlReconnectStatus = __MIDL_IMsRdpClient8_0001; 
  RemoteActionType = __MIDL_IRemoteDesktopClientActions_0001; 
  SnapshotEncodingType = __MIDL_IRemoteDesktopClientActions_0002; 
  SnapshotFormatType = __MIDL_IRemoteDesktopClientActions_0003; 

// *********************************************************************//
// DispIntf:  IMsTscAxEvents
// Flags:     (4096) Dispatchable
// GUID:      {336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}
// *********************************************************************//
  IMsTscAxEvents = dispinterface
    ['{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}']
    procedure OnConnecting; dispid 1;
    procedure OnConnected; dispid 2;
    procedure OnLoginComplete; dispid 3;
    procedure OnDisconnected(discReason: Integer); dispid 4;
    procedure OnEnterFullScreenMode; dispid 5;
    procedure OnLeaveFullScreenMode; dispid 6;
    procedure OnChannelReceivedData(const chanName: WideString; const data: WideString); dispid 7;
    procedure OnRequestGoFullScreen; dispid 8;
    procedure OnRequestLeaveFullScreen; dispid 9;
    procedure OnFatalError(errorCode: Integer); dispid 10;
    procedure OnWarning(warningCode: Integer); dispid 11;
    procedure OnRemoteDesktopSizeChange(width: Integer; height: Integer); dispid 12;
    procedure OnIdleTimeoutNotification; dispid 13;
    procedure OnRequestContainerMinimize; dispid 14;
    function OnConfirmClose: WordBool; dispid 15;
    function OnReceivedTSPublicKey(const publicKey: WideString): WordBool; dispid 16;
    function OnAutoReconnecting(disconnectReason: Integer; attemptCount: Integer): AutoReconnectContinueState; dispid 17;
    procedure OnAuthenticationWarningDisplayed; dispid 18;
    procedure OnAuthenticationWarningDismissed; dispid 19;
    procedure OnRemoteProgramResult(const bstrRemoteProgram: WideString; 
                                    lError: RemoteProgramResult; vbIsExecutable: WordBool); dispid 20;
    procedure OnRemoteProgramDisplayed(vbDisplayed: WordBool; uDisplayInformation: LongWord); dispid 21;
    procedure OnRemoteWindowDisplayed(vbDisplayed: WordBool; 
                                      var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                      windowAttribute: RemoteWindowDisplayedAttribute); dispid 29;
    procedure OnLogonError(lError: Integer); dispid 22;
    procedure OnFocusReleased(iDirection: SYSINT); dispid 23;
    procedure OnUserNameAcquired(const bstrUserName: WideString); dispid 24;
    procedure OnMouseInputModeChanged(fMouseModeRelative: WordBool); dispid 26;
    procedure OnServiceMessageReceived(const serviceMessage: WideString); dispid 28;
    procedure OnConnectionBarPullDown; dispid 30;
    procedure OnNetworkStatusChanged(qualityLevel: LongWord; bandwidth: Integer; rtt: Integer); dispid 32;
    procedure OnDevicesButtonPressed; dispid 35;
    procedure OnAutoReconnected; dispid 33;
    procedure OnAutoReconnecting2(disconnectReason: Integer; networkAvailable: WordBool; 
                                  attemptCount: Integer; maxAttemptCount: Integer); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsTscAx_Redist
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {327BB5CD-834E-4400-AEF2-B30E15E5D682}
// *********************************************************************//
  IMsTscAx_Redist = interface(IDispatch)
    ['{327BB5CD-834E-4400-AEF2-B30E15E5D682}']
  end;

// *********************************************************************//
// DispIntf:  IMsTscAx_RedistDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {327BB5CD-834E-4400-AEF2-B30E15E5D682}
// *********************************************************************//
  IMsTscAx_RedistDisp = dispinterface
    ['{327BB5CD-834E-4400-AEF2-B30E15E5D682}']
  end;

// *********************************************************************//
// Interface: IMsTscAx
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {8C11EFAE-92C3-11D1-BC1E-00C04FA31489}
// *********************************************************************//
  IMsTscAx = interface(IMsTscAx_Redist)
    ['{8C11EFAE-92C3-11D1-BC1E-00C04FA31489}']
    procedure Set_Server(const pServer: WideString); safecall;
    function Get_Server: WideString; safecall;
    procedure Set_Domain(const pDomain: WideString); safecall;
    function Get_Domain: WideString; safecall;
    procedure Set_UserName(const pUserName: WideString); safecall;
    function Get_UserName: WideString; safecall;
    procedure Set_DisconnectedText(const pDisconnectedText: WideString); safecall;
    function Get_DisconnectedText: WideString; safecall;
    procedure Set_ConnectingText(const pConnectingText: WideString); safecall;
    function Get_ConnectingText: WideString; safecall;
    function Get_Connected: Smallint; safecall;
    procedure Set_DesktopWidth(pVal: Integer); safecall;
    function Get_DesktopWidth: Integer; safecall;
    procedure Set_DesktopHeight(pVal: Integer); safecall;
    function Get_DesktopHeight: Integer; safecall;
    procedure Set_StartConnected(pfStartConnected: Integer); safecall;
    function Get_StartConnected: Integer; safecall;
    function Get_HorizontalScrollBarVisible: Integer; safecall;
    function Get_VerticalScrollBarVisible: Integer; safecall;
    procedure Set_FullScreenTitle(const Param1: WideString); safecall;
    function Get_CipherStrength: Integer; safecall;
    function Get_Version: WideString; safecall;
    function Get_SecuredSettingsEnabled: Integer; safecall;
    function Get_SecuredSettings: IMsTscSecuredSettings; safecall;
    function Get_AdvancedSettings: IMsTscAdvancedSettings; safecall;
    function Get_Debugger: IMsTscDebug; safecall;
    procedure Connect; safecall;
    procedure Disconnect; safecall;
    procedure CreateVirtualChannels(const newVal: WideString); safecall;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); safecall;
    property Server: WideString read Get_Server write Set_Server;
    property Domain: WideString read Get_Domain write Set_Domain;
    property UserName: WideString read Get_UserName write Set_UserName;
    property DisconnectedText: WideString read Get_DisconnectedText write Set_DisconnectedText;
    property ConnectingText: WideString read Get_ConnectingText write Set_ConnectingText;
    property Connected: Smallint read Get_Connected;
    property DesktopWidth: Integer read Get_DesktopWidth write Set_DesktopWidth;
    property DesktopHeight: Integer read Get_DesktopHeight write Set_DesktopHeight;
    property StartConnected: Integer read Get_StartConnected write Set_StartConnected;
    property HorizontalScrollBarVisible: Integer read Get_HorizontalScrollBarVisible;
    property VerticalScrollBarVisible: Integer read Get_VerticalScrollBarVisible;
    property FullScreenTitle: WideString write Set_FullScreenTitle;
    property CipherStrength: Integer read Get_CipherStrength;
    property Version: WideString read Get_Version;
    property SecuredSettingsEnabled: Integer read Get_SecuredSettingsEnabled;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
  end;

// *********************************************************************//
// DispIntf:  IMsTscAxDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {8C11EFAE-92C3-11D1-BC1E-00C04FA31489}
// *********************************************************************//
  IMsTscAxDisp = dispinterface
    ['{8C11EFAE-92C3-11D1-BC1E-00C04FA31489}']
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClient
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {92B4A539-7115-4B7C-A5A9-E5D9EFC2780A}
// *********************************************************************//
  IMsRdpClient = interface(IMsTscAx)
    ['{92B4A539-7115-4B7C-A5A9-E5D9EFC2780A}']
    procedure Set_ColorDepth(pcolorDepth: Integer); safecall;
    function Get_ColorDepth: Integer; safecall;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings; safecall;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings; safecall;
    function Get_ExtendedDisconnectReason: ExtendedDisconnectReasonCode; safecall;
    procedure Set_FullScreen(pfFullScreen: WordBool); safecall;
    function Get_FullScreen: WordBool; safecall;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); safecall;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; safecall;
    function RequestClose: ControlCloseStatus; safecall;
    property ColorDepth: Integer read Get_ColorDepth write Set_ColorDepth;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode read Get_ExtendedDisconnectReason;
    property FullScreen: WordBool read Get_FullScreen write Set_FullScreen;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {92B4A539-7115-4B7C-A5A9-E5D9EFC2780A}
// *********************************************************************//
  IMsRdpClientDisp = dispinterface
    ['{92B4A539-7115-4B7C-A5A9-E5D9EFC2780A}']
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsTscNonScriptable
// Flags:     (0)
// GUID:      {C1E6743A-41C1-4A74-832A-0DD06C1C7A0E}
// *********************************************************************//
  IMsTscNonScriptable = interface(IUnknown)
    ['{C1E6743A-41C1-4A74-832A-0DD06C1C7A0E}']
    function Set_ClearTextPassword(const Param1: WideString): HResult; stdcall;
    function Set_PortablePassword(const pPortablePass: WideString): HResult; stdcall;
    function Get_PortablePassword(out pPortablePass: WideString): HResult; stdcall;
    function Set_PortableSalt(const pPortableSalt: WideString): HResult; stdcall;
    function Get_PortableSalt(out pPortableSalt: WideString): HResult; stdcall;
    function Set_BinaryPassword(const pBinaryPassword: WideString): HResult; stdcall;
    function Get_BinaryPassword(out pBinaryPassword: WideString): HResult; stdcall;
    function Set_BinarySalt(const pSalt: WideString): HResult; stdcall;
    function Get_BinarySalt(out pSalt: WideString): HResult; stdcall;
    function ResetPassword: HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsTscSecuredSettings
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {C9D65442-A0F9-45B2-8F73-D61D2DB8CBB6}
// *********************************************************************//
  IMsTscSecuredSettings = interface(IDispatch)
    ['{C9D65442-A0F9-45B2-8F73-D61D2DB8CBB6}']
    procedure Set_StartProgram(const pStartProgram: WideString); safecall;
    function Get_StartProgram: WideString; safecall;
    procedure Set_WorkDir(const pWorkDir: WideString); safecall;
    function Get_WorkDir: WideString; safecall;
    procedure Set_FullScreen(pfFullScreen: Integer); safecall;
    function Get_FullScreen: Integer; safecall;
    property StartProgram: WideString read Get_StartProgram write Set_StartProgram;
    property WorkDir: WideString read Get_WorkDir write Set_WorkDir;
    property FullScreen: Integer read Get_FullScreen write Set_FullScreen;
  end;

// *********************************************************************//
// DispIntf:  IMsTscSecuredSettingsDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {C9D65442-A0F9-45B2-8F73-D61D2DB8CBB6}
// *********************************************************************//
  IMsTscSecuredSettingsDisp = dispinterface
    ['{C9D65442-A0F9-45B2-8F73-D61D2DB8CBB6}']
    property StartProgram: WideString dispid 1;
    property WorkDir: WideString dispid 2;
    property FullScreen: Integer dispid 3;
  end;

// *********************************************************************//
// Interface: IMsTscAdvancedSettings
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {809945CC-4B3B-4A92-A6B0-DBF9B5F2EF2D}
// *********************************************************************//
  IMsTscAdvancedSettings = interface(IDispatch)
    ['{809945CC-4B3B-4A92-A6B0-DBF9B5F2EF2D}']
    procedure Set_Compress(pcompress: Integer); safecall;
    function Get_Compress: Integer; safecall;
    procedure Set_BitmapPeristence(pbitmapPeristence: Integer); safecall;
    function Get_BitmapPeristence: Integer; safecall;
    procedure Set_allowBackgroundInput(pallowBackgroundInput: Integer); safecall;
    function Get_allowBackgroundInput: Integer; safecall;
    procedure Set_KeyBoardLayoutStr(const Param1: WideString); safecall;
    procedure Set_PluginDlls(const Param1: WideString); safecall;
    procedure Set_IconFile(const Param1: WideString); safecall;
    procedure Set_IconIndex(Param1: Integer); safecall;
    procedure Set_ContainerHandledFullScreen(pContainerHandledFullScreen: Integer); safecall;
    function Get_ContainerHandledFullScreen: Integer; safecall;
    procedure Set_DisableRdpdr(pDisableRdpdr: Integer); safecall;
    function Get_DisableRdpdr: Integer; safecall;
    property Compress: Integer read Get_Compress write Set_Compress;
    property BitmapPeristence: Integer read Get_BitmapPeristence write Set_BitmapPeristence;
    property allowBackgroundInput: Integer read Get_allowBackgroundInput write Set_allowBackgroundInput;
    property KeyBoardLayoutStr: WideString write Set_KeyBoardLayoutStr;
    property PluginDlls: WideString write Set_PluginDlls;
    property IconFile: WideString write Set_IconFile;
    property IconIndex: Integer write Set_IconIndex;
    property ContainerHandledFullScreen: Integer read Get_ContainerHandledFullScreen write Set_ContainerHandledFullScreen;
    property DisableRdpdr: Integer read Get_DisableRdpdr write Set_DisableRdpdr;
  end;

// *********************************************************************//
// DispIntf:  IMsTscAdvancedSettingsDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {809945CC-4B3B-4A92-A6B0-DBF9B5F2EF2D}
// *********************************************************************//
  IMsTscAdvancedSettingsDisp = dispinterface
    ['{809945CC-4B3B-4A92-A6B0-DBF9B5F2EF2D}']
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: IMsTscDebug
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {209D0EB9-6254-47B1-9033-A98DAE55BB27}
// *********************************************************************//
  IMsTscDebug = interface(IDispatch)
    ['{209D0EB9-6254-47B1-9033-A98DAE55BB27}']
    procedure Set_HatchBitmapPDU(phatchBitmapPDU: Integer); safecall;
    function Get_HatchBitmapPDU: Integer; safecall;
    procedure Set_HatchSSBOrder(phatchSSBOrder: Integer); safecall;
    function Get_HatchSSBOrder: Integer; safecall;
    procedure Set_HatchMembltOrder(phatchMembltOrder: Integer); safecall;
    function Get_HatchMembltOrder: Integer; safecall;
    procedure Set_HatchIndexPDU(phatchIndexPDU: Integer); safecall;
    function Get_HatchIndexPDU: Integer; safecall;
    procedure Set_LabelMemblt(plabelMemblt: Integer); safecall;
    function Get_LabelMemblt: Integer; safecall;
    procedure Set_BitmapCacheMonitor(pbitmapCacheMonitor: Integer); safecall;
    function Get_BitmapCacheMonitor: Integer; safecall;
    procedure Set_MallocFailuresPercent(pmallocFailuresPercent: Integer); safecall;
    function Get_MallocFailuresPercent: Integer; safecall;
    procedure Set_MallocHugeFailuresPercent(pmallocHugeFailuresPercent: Integer); safecall;
    function Get_MallocHugeFailuresPercent: Integer; safecall;
    procedure Set_NetThroughput(NetThroughput: Integer); safecall;
    function Get_NetThroughput: Integer; safecall;
    procedure Set_CLXCmdLine(const pCLXCmdLine: WideString); safecall;
    function Get_CLXCmdLine: WideString; safecall;
    procedure Set_CLXDll(const pCLXDll: WideString); safecall;
    function Get_CLXDll: WideString; safecall;
    procedure Set_RemoteProgramsHatchVisibleRegion(pcbHatch: Integer); safecall;
    function Get_RemoteProgramsHatchVisibleRegion: Integer; safecall;
    procedure Set_RemoteProgramsHatchVisibleNoDataRegion(pcbHatch: Integer); safecall;
    function Get_RemoteProgramsHatchVisibleNoDataRegion: Integer; safecall;
    procedure Set_RemoteProgramsHatchNonVisibleRegion(pcbHatch: Integer); safecall;
    function Get_RemoteProgramsHatchNonVisibleRegion: Integer; safecall;
    procedure Set_RemoteProgramsHatchWindow(pcbHatch: Integer); safecall;
    function Get_RemoteProgramsHatchWindow: Integer; safecall;
    procedure Set_RemoteProgramsStayConnectOnBadCaps(pcbStayConnected: Integer); safecall;
    function Get_RemoteProgramsStayConnectOnBadCaps: Integer; safecall;
    function Get_ControlType: SYSUINT; safecall;
    procedure Set_DecodeGfx(Param1: WordBool); safecall;
    property HatchBitmapPDU: Integer read Get_HatchBitmapPDU write Set_HatchBitmapPDU;
    property HatchSSBOrder: Integer read Get_HatchSSBOrder write Set_HatchSSBOrder;
    property HatchMembltOrder: Integer read Get_HatchMembltOrder write Set_HatchMembltOrder;
    property HatchIndexPDU: Integer read Get_HatchIndexPDU write Set_HatchIndexPDU;
    property LabelMemblt: Integer read Get_LabelMemblt write Set_LabelMemblt;
    property BitmapCacheMonitor: Integer read Get_BitmapCacheMonitor write Set_BitmapCacheMonitor;
    property MallocFailuresPercent: Integer read Get_MallocFailuresPercent write Set_MallocFailuresPercent;
    property MallocHugeFailuresPercent: Integer read Get_MallocHugeFailuresPercent write Set_MallocHugeFailuresPercent;
    property NetThroughput: Integer read Get_NetThroughput write Set_NetThroughput;
    property CLXCmdLine: WideString read Get_CLXCmdLine write Set_CLXCmdLine;
    property CLXDll: WideString read Get_CLXDll write Set_CLXDll;
    property RemoteProgramsHatchVisibleRegion: Integer read Get_RemoteProgramsHatchVisibleRegion write Set_RemoteProgramsHatchVisibleRegion;
    property RemoteProgramsHatchVisibleNoDataRegion: Integer read Get_RemoteProgramsHatchVisibleNoDataRegion write Set_RemoteProgramsHatchVisibleNoDataRegion;
    property RemoteProgramsHatchNonVisibleRegion: Integer read Get_RemoteProgramsHatchNonVisibleRegion write Set_RemoteProgramsHatchNonVisibleRegion;
    property RemoteProgramsHatchWindow: Integer read Get_RemoteProgramsHatchWindow write Set_RemoteProgramsHatchWindow;
    property RemoteProgramsStayConnectOnBadCaps: Integer read Get_RemoteProgramsStayConnectOnBadCaps write Set_RemoteProgramsStayConnectOnBadCaps;
    property ControlType: SYSUINT read Get_ControlType;
    property DecodeGfx: WordBool write Set_DecodeGfx;
  end;

// *********************************************************************//
// DispIntf:  IMsTscDebugDisp
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {209D0EB9-6254-47B1-9033-A98DAE55BB27}
// *********************************************************************//
  IMsTscDebugDisp = dispinterface
    ['{209D0EB9-6254-47B1-9033-A98DAE55BB27}']
    property HatchBitmapPDU: Integer dispid 200;
    property HatchSSBOrder: Integer dispid 201;
    property HatchMembltOrder: Integer dispid 202;
    property HatchIndexPDU: Integer dispid 203;
    property LabelMemblt: Integer dispid 204;
    property BitmapCacheMonitor: Integer dispid 205;
    property MallocFailuresPercent: Integer dispid 206;
    property MallocHugeFailuresPercent: Integer dispid 207;
    property NetThroughput: Integer dispid 208;
    property CLXCmdLine: WideString dispid 209;
    property CLXDll: WideString dispid 210;
    property RemoteProgramsHatchVisibleRegion: Integer dispid 211;
    property RemoteProgramsHatchVisibleNoDataRegion: Integer dispid 212;
    property RemoteProgramsHatchNonVisibleRegion: Integer dispid 213;
    property RemoteProgramsHatchWindow: Integer dispid 214;
    property RemoteProgramsStayConnectOnBadCaps: Integer dispid 215;
    property ControlType: SYSUINT readonly dispid 216;
    property DecodeGfx: WordBool writeonly dispid 217;
  end;

// *********************************************************************//
// Interface: IMsRdpClientAdvancedSettings
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {3C65B4AB-12B3-465B-ACD4-B8DAD3BFF9E2}
// *********************************************************************//
  IMsRdpClientAdvancedSettings = interface(IMsTscAdvancedSettings)
    ['{3C65B4AB-12B3-465B-ACD4-B8DAD3BFF9E2}']
    procedure Set_SmoothScroll(psmoothScroll: Integer); safecall;
    function Get_SmoothScroll: Integer; safecall;
    procedure Set_AcceleratorPassthrough(pacceleratorPassthrough: Integer); safecall;
    function Get_AcceleratorPassthrough: Integer; safecall;
    procedure Set_ShadowBitmap(pshadowBitmap: Integer); safecall;
    function Get_ShadowBitmap: Integer; safecall;
    procedure Set_TransportType(ptransportType: Integer); safecall;
    function Get_TransportType: Integer; safecall;
    procedure Set_SasSequence(psasSequence: Integer); safecall;
    function Get_SasSequence: Integer; safecall;
    procedure Set_EncryptionEnabled(pencryptionEnabled: Integer); safecall;
    function Get_EncryptionEnabled: Integer; safecall;
    procedure Set_DedicatedTerminal(pdedicatedTerminal: Integer); safecall;
    function Get_DedicatedTerminal: Integer; safecall;
    procedure Set_RDPPort(prdpPort: Integer); safecall;
    function Get_RDPPort: Integer; safecall;
    procedure Set_EnableMouse(penableMouse: Integer); safecall;
    function Get_EnableMouse: Integer; safecall;
    procedure Set_DisableCtrlAltDel(pdisableCtrlAltDel: Integer); safecall;
    function Get_DisableCtrlAltDel: Integer; safecall;
    procedure Set_EnableWindowsKey(penableWindowsKey: Integer); safecall;
    function Get_EnableWindowsKey: Integer; safecall;
    procedure Set_DoubleClickDetect(pdoubleClickDetect: Integer); safecall;
    function Get_DoubleClickDetect: Integer; safecall;
    procedure Set_MaximizeShell(pmaximizeShell: Integer); safecall;
    function Get_MaximizeShell: Integer; safecall;
    procedure Set_HotKeyFullScreen(photKeyFullScreen: Integer); safecall;
    function Get_HotKeyFullScreen: Integer; safecall;
    procedure Set_HotKeyCtrlEsc(photKeyCtrlEsc: Integer); safecall;
    function Get_HotKeyCtrlEsc: Integer; safecall;
    procedure Set_HotKeyAltEsc(photKeyAltEsc: Integer); safecall;
    function Get_HotKeyAltEsc: Integer; safecall;
    procedure Set_HotKeyAltTab(photKeyAltTab: Integer); safecall;
    function Get_HotKeyAltTab: Integer; safecall;
    procedure Set_HotKeyAltShiftTab(photKeyAltShiftTab: Integer); safecall;
    function Get_HotKeyAltShiftTab: Integer; safecall;
    procedure Set_HotKeyAltSpace(photKeyAltSpace: Integer); safecall;
    function Get_HotKeyAltSpace: Integer; safecall;
    procedure Set_HotKeyCtrlAltDel(photKeyCtrlAltDel: Integer); safecall;
    function Get_HotKeyCtrlAltDel: Integer; safecall;
    procedure Set_orderDrawThreshold(porderDrawThreshold: Integer); safecall;
    function Get_orderDrawThreshold: Integer; safecall;
    procedure Set_BitmapCacheSize(pbitmapCacheSize: Integer); safecall;
    function Get_BitmapCacheSize: Integer; safecall;
    procedure Set_BitmapVirtualCacheSize(pbitmapVirtualCacheSize: Integer); safecall;
    function Get_BitmapVirtualCacheSize: Integer; safecall;
    procedure Set_ScaleBitmapCachesByBPP(pbScale: Integer); safecall;
    function Get_ScaleBitmapCachesByBPP: Integer; safecall;
    procedure Set_NumBitmapCaches(pnumBitmapCaches: Integer); safecall;
    function Get_NumBitmapCaches: Integer; safecall;
    procedure Set_CachePersistenceActive(pcachePersistenceActive: Integer); safecall;
    function Get_CachePersistenceActive: Integer; safecall;
    procedure Set_PersistCacheDirectory(const Param1: WideString); safecall;
    procedure Set_brushSupportLevel(pbrushSupportLevel: Integer); safecall;
    function Get_brushSupportLevel: Integer; safecall;
    procedure Set_minInputSendInterval(pminInputSendInterval: Integer); safecall;
    function Get_minInputSendInterval: Integer; safecall;
    procedure Set_InputEventsAtOnce(pinputEventsAtOnce: Integer); safecall;
    function Get_InputEventsAtOnce: Integer; safecall;
    procedure Set_maxEventCount(pmaxEventCount: Integer); safecall;
    function Get_maxEventCount: Integer; safecall;
    procedure Set_keepAliveInterval(pkeepAliveInterval: Integer); safecall;
    function Get_keepAliveInterval: Integer; safecall;
    procedure Set_shutdownTimeout(pshutdownTimeout: Integer); safecall;
    function Get_shutdownTimeout: Integer; safecall;
    procedure Set_overallConnectionTimeout(poverallConnectionTimeout: Integer); safecall;
    function Get_overallConnectionTimeout: Integer; safecall;
    procedure Set_singleConnectionTimeout(psingleConnectionTimeout: Integer); safecall;
    function Get_singleConnectionTimeout: Integer; safecall;
    procedure Set_KeyboardType(pkeyboardType: Integer); safecall;
    function Get_KeyboardType: Integer; safecall;
    procedure Set_KeyboardSubType(pkeyboardSubType: Integer); safecall;
    function Get_KeyboardSubType: Integer; safecall;
    procedure Set_KeyboardFunctionKey(pkeyboardFunctionKey: Integer); safecall;
    function Get_KeyboardFunctionKey: Integer; safecall;
    procedure Set_WinceFixedPalette(pwinceFixedPalette: Integer); safecall;
    function Get_WinceFixedPalette: Integer; safecall;
    procedure Set_ConnectToServerConsole(pConnectToConsole: WordBool); safecall;
    function Get_ConnectToServerConsole: WordBool; safecall;
    procedure Set_BitmapPersistence(pbitmapPersistence: Integer); safecall;
    function Get_BitmapPersistence: Integer; safecall;
    procedure Set_MinutesToIdleTimeout(pminutesToIdleTimeout: Integer); safecall;
    function Get_MinutesToIdleTimeout: Integer; safecall;
    procedure Set_SmartSizing(pfSmartSizing: WordBool); safecall;
    function Get_SmartSizing: WordBool; safecall;
    procedure Set_RdpdrLocalPrintingDocName(const pLocalPrintingDocName: WideString); safecall;
    function Get_RdpdrLocalPrintingDocName: WideString; safecall;
    procedure Set_RdpdrClipCleanTempDirString(const clipCleanTempDirString: WideString); safecall;
    function Get_RdpdrClipCleanTempDirString: WideString; safecall;
    procedure Set_RdpdrClipPasteInfoString(const clipPasteInfoString: WideString); safecall;
    function Get_RdpdrClipPasteInfoString: WideString; safecall;
    procedure Set_ClearTextPassword(const Param1: WideString); safecall;
    procedure Set_DisplayConnectionBar(pDisplayConnectionBar: WordBool); safecall;
    function Get_DisplayConnectionBar: WordBool; safecall;
    procedure Set_PinConnectionBar(pPinConnectionBar: WordBool); safecall;
    function Get_PinConnectionBar: WordBool; safecall;
    procedure Set_GrabFocusOnConnect(pfGrabFocusOnConnect: WordBool); safecall;
    function Get_GrabFocusOnConnect: WordBool; safecall;
    procedure Set_LoadBalanceInfo(const pLBInfo: WideString); safecall;
    function Get_LoadBalanceInfo: WideString; safecall;
    procedure Set_RedirectDrives(pRedirectDrives: WordBool); safecall;
    function Get_RedirectDrives: WordBool; safecall;
    procedure Set_RedirectPrinters(pRedirectPrinters: WordBool); safecall;
    function Get_RedirectPrinters: WordBool; safecall;
    procedure Set_RedirectPorts(pRedirectPorts: WordBool); safecall;
    function Get_RedirectPorts: WordBool; safecall;
    procedure Set_RedirectSmartCards(pRedirectSmartCards: WordBool); safecall;
    function Get_RedirectSmartCards: WordBool; safecall;
    procedure Set_BitmapVirtualCache16BppSize(pBitmapVirtualCache16BppSize: Integer); safecall;
    function Get_BitmapVirtualCache16BppSize: Integer; safecall;
    procedure Set_BitmapVirtualCache24BppSize(pBitmapVirtualCache24BppSize: Integer); safecall;
    function Get_BitmapVirtualCache24BppSize: Integer; safecall;
    procedure Set_PerformanceFlags(pDisableList: Integer); safecall;
    function Get_PerformanceFlags: Integer; safecall;
    procedure Set_ConnectWithEndpoint(const Param1: OleVariant); safecall;
    procedure Set_NotifyTSPublicKey(pfNotify: WordBool); safecall;
    function Get_NotifyTSPublicKey: WordBool; safecall;
    property SmoothScroll: Integer read Get_SmoothScroll write Set_SmoothScroll;
    property AcceleratorPassthrough: Integer read Get_AcceleratorPassthrough write Set_AcceleratorPassthrough;
    property ShadowBitmap: Integer read Get_ShadowBitmap write Set_ShadowBitmap;
    property TransportType: Integer read Get_TransportType write Set_TransportType;
    property SasSequence: Integer read Get_SasSequence write Set_SasSequence;
    property EncryptionEnabled: Integer read Get_EncryptionEnabled write Set_EncryptionEnabled;
    property DedicatedTerminal: Integer read Get_DedicatedTerminal write Set_DedicatedTerminal;
    property RDPPort: Integer read Get_RDPPort write Set_RDPPort;
    property EnableMouse: Integer read Get_EnableMouse write Set_EnableMouse;
    property DisableCtrlAltDel: Integer read Get_DisableCtrlAltDel write Set_DisableCtrlAltDel;
    property EnableWindowsKey: Integer read Get_EnableWindowsKey write Set_EnableWindowsKey;
    property DoubleClickDetect: Integer read Get_DoubleClickDetect write Set_DoubleClickDetect;
    property MaximizeShell: Integer read Get_MaximizeShell write Set_MaximizeShell;
    property HotKeyFullScreen: Integer read Get_HotKeyFullScreen write Set_HotKeyFullScreen;
    property HotKeyCtrlEsc: Integer read Get_HotKeyCtrlEsc write Set_HotKeyCtrlEsc;
    property HotKeyAltEsc: Integer read Get_HotKeyAltEsc write Set_HotKeyAltEsc;
    property HotKeyAltTab: Integer read Get_HotKeyAltTab write Set_HotKeyAltTab;
    property HotKeyAltShiftTab: Integer read Get_HotKeyAltShiftTab write Set_HotKeyAltShiftTab;
    property HotKeyAltSpace: Integer read Get_HotKeyAltSpace write Set_HotKeyAltSpace;
    property HotKeyCtrlAltDel: Integer read Get_HotKeyCtrlAltDel write Set_HotKeyCtrlAltDel;
    property orderDrawThreshold: Integer read Get_orderDrawThreshold write Set_orderDrawThreshold;
    property BitmapCacheSize: Integer read Get_BitmapCacheSize write Set_BitmapCacheSize;
    property BitmapVirtualCacheSize: Integer read Get_BitmapVirtualCacheSize write Set_BitmapVirtualCacheSize;
    property ScaleBitmapCachesByBPP: Integer read Get_ScaleBitmapCachesByBPP write Set_ScaleBitmapCachesByBPP;
    property NumBitmapCaches: Integer read Get_NumBitmapCaches write Set_NumBitmapCaches;
    property CachePersistenceActive: Integer read Get_CachePersistenceActive write Set_CachePersistenceActive;
    property PersistCacheDirectory: WideString write Set_PersistCacheDirectory;
    property brushSupportLevel: Integer read Get_brushSupportLevel write Set_brushSupportLevel;
    property minInputSendInterval: Integer read Get_minInputSendInterval write Set_minInputSendInterval;
    property InputEventsAtOnce: Integer read Get_InputEventsAtOnce write Set_InputEventsAtOnce;
    property maxEventCount: Integer read Get_maxEventCount write Set_maxEventCount;
    property keepAliveInterval: Integer read Get_keepAliveInterval write Set_keepAliveInterval;
    property shutdownTimeout: Integer read Get_shutdownTimeout write Set_shutdownTimeout;
    property overallConnectionTimeout: Integer read Get_overallConnectionTimeout write Set_overallConnectionTimeout;
    property singleConnectionTimeout: Integer read Get_singleConnectionTimeout write Set_singleConnectionTimeout;
    property KeyboardType: Integer read Get_KeyboardType write Set_KeyboardType;
    property KeyboardSubType: Integer read Get_KeyboardSubType write Set_KeyboardSubType;
    property KeyboardFunctionKey: Integer read Get_KeyboardFunctionKey write Set_KeyboardFunctionKey;
    property WinceFixedPalette: Integer read Get_WinceFixedPalette write Set_WinceFixedPalette;
    property ConnectToServerConsole: WordBool read Get_ConnectToServerConsole write Set_ConnectToServerConsole;
    property BitmapPersistence: Integer read Get_BitmapPersistence write Set_BitmapPersistence;
    property MinutesToIdleTimeout: Integer read Get_MinutesToIdleTimeout write Set_MinutesToIdleTimeout;
    property SmartSizing: WordBool read Get_SmartSizing write Set_SmartSizing;
    property RdpdrLocalPrintingDocName: WideString read Get_RdpdrLocalPrintingDocName write Set_RdpdrLocalPrintingDocName;
    property RdpdrClipCleanTempDirString: WideString read Get_RdpdrClipCleanTempDirString write Set_RdpdrClipCleanTempDirString;
    property RdpdrClipPasteInfoString: WideString read Get_RdpdrClipPasteInfoString write Set_RdpdrClipPasteInfoString;
    property ClearTextPassword: WideString write Set_ClearTextPassword;
    property DisplayConnectionBar: WordBool read Get_DisplayConnectionBar write Set_DisplayConnectionBar;
    property PinConnectionBar: WordBool read Get_PinConnectionBar write Set_PinConnectionBar;
    property GrabFocusOnConnect: WordBool read Get_GrabFocusOnConnect write Set_GrabFocusOnConnect;
    property LoadBalanceInfo: WideString read Get_LoadBalanceInfo write Set_LoadBalanceInfo;
    property RedirectDrives: WordBool read Get_RedirectDrives write Set_RedirectDrives;
    property RedirectPrinters: WordBool read Get_RedirectPrinters write Set_RedirectPrinters;
    property RedirectPorts: WordBool read Get_RedirectPorts write Set_RedirectPorts;
    property RedirectSmartCards: WordBool read Get_RedirectSmartCards write Set_RedirectSmartCards;
    property BitmapVirtualCache16BppSize: Integer read Get_BitmapVirtualCache16BppSize write Set_BitmapVirtualCache16BppSize;
    property BitmapVirtualCache24BppSize: Integer read Get_BitmapVirtualCache24BppSize write Set_BitmapVirtualCache24BppSize;
    property PerformanceFlags: Integer read Get_PerformanceFlags write Set_PerformanceFlags;
    //property ConnectWithEndpoint: POleVariant1 write Set_ConnectWithEndpoint;
    property NotifyTSPublicKey: WordBool read Get_NotifyTSPublicKey write Set_NotifyTSPublicKey;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientAdvancedSettingsDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {3C65B4AB-12B3-465B-ACD4-B8DAD3BFF9E2}
// *********************************************************************//
  IMsRdpClientAdvancedSettingsDisp = dispinterface
    ['{3C65B4AB-12B3-465B-ACD4-B8DAD3BFF9E2}']
    property SmoothScroll: Integer dispid 101;
    property AcceleratorPassthrough: Integer dispid 102;
    property ShadowBitmap: Integer dispid 103;
    property TransportType: Integer dispid 104;
    property SasSequence: Integer dispid 105;
    property EncryptionEnabled: Integer dispid 106;
    property DedicatedTerminal: Integer dispid 107;
    property RDPPort: Integer dispid 108;
    property EnableMouse: Integer dispid 109;
    property DisableCtrlAltDel: Integer dispid 110;
    property EnableWindowsKey: Integer dispid 111;
    property DoubleClickDetect: Integer dispid 112;
    property MaximizeShell: Integer dispid 113;
    property HotKeyFullScreen: Integer dispid 114;
    property HotKeyCtrlEsc: Integer dispid 115;
    property HotKeyAltEsc: Integer dispid 116;
    property HotKeyAltTab: Integer dispid 117;
    property HotKeyAltShiftTab: Integer dispid 118;
    property HotKeyAltSpace: Integer dispid 119;
    property HotKeyCtrlAltDel: Integer dispid 120;
    property orderDrawThreshold: Integer dispid 123;
    property BitmapCacheSize: Integer dispid 124;
    property BitmapVirtualCacheSize: Integer dispid 125;
    property ScaleBitmapCachesByBPP: Integer dispid 175;
    property NumBitmapCaches: Integer dispid 126;
    property CachePersistenceActive: Integer dispid 127;
    property PersistCacheDirectory: WideString writeonly dispid 138;
    property brushSupportLevel: Integer dispid 156;
    property minInputSendInterval: Integer dispid 157;
    property InputEventsAtOnce: Integer dispid 158;
    property maxEventCount: Integer dispid 159;
    property keepAliveInterval: Integer dispid 160;
    property shutdownTimeout: Integer dispid 163;
    property overallConnectionTimeout: Integer dispid 164;
    property singleConnectionTimeout: Integer dispid 165;
    property KeyboardType: Integer dispid 166;
    property KeyboardSubType: Integer dispid 167;
    property KeyboardFunctionKey: Integer dispid 168;
    property WinceFixedPalette: Integer dispid 169;
    property ConnectToServerConsole: WordBool dispid 178;
    property BitmapPersistence: Integer dispid 182;
    property MinutesToIdleTimeout: Integer dispid 183;
    property SmartSizing: WordBool dispid 184;
    property RdpdrLocalPrintingDocName: WideString dispid 185;
    property RdpdrClipCleanTempDirString: WideString dispid 201;
    property RdpdrClipPasteInfoString: WideString dispid 202;
    property ClearTextPassword: WideString writeonly dispid 186;
    property DisplayConnectionBar: WordBool dispid 187;
    property PinConnectionBar: WordBool dispid 188;
    property GrabFocusOnConnect: WordBool dispid 189;
    property LoadBalanceInfo: WideString dispid 190;
    property RedirectDrives: WordBool dispid 191;
    property RedirectPrinters: WordBool dispid 192;
    property RedirectPorts: WordBool dispid 193;
    property RedirectSmartCards: WordBool dispid 194;
    property BitmapVirtualCache16BppSize: Integer dispid 195;
    property BitmapVirtualCache24BppSize: Integer dispid 196;
    property PerformanceFlags: Integer dispid 200;
    property ConnectWithEndpoint: {NOT_OLEAUTO(POleVariant1)}OleVariant writeonly dispid 203;
    property NotifyTSPublicKey: WordBool dispid 204;
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: IMsRdpClientSecuredSettings
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {605BEFCF-39C1-45CC-A811-068FB7BE346D}
// *********************************************************************//
  IMsRdpClientSecuredSettings = interface(IMsTscSecuredSettings)
    ['{605BEFCF-39C1-45CC-A811-068FB7BE346D}']
    procedure Set_KeyboardHookMode(pkeyboardHookMode: Integer); safecall;
    function Get_KeyboardHookMode: Integer; safecall;
    procedure Set_AudioRedirectionMode(pAudioRedirectionMode: Integer); safecall;
    function Get_AudioRedirectionMode: Integer; safecall;
    property KeyboardHookMode: Integer read Get_KeyboardHookMode write Set_KeyboardHookMode;
    property AudioRedirectionMode: Integer read Get_AudioRedirectionMode write Set_AudioRedirectionMode;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientSecuredSettingsDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {605BEFCF-39C1-45CC-A811-068FB7BE346D}
// *********************************************************************//
  IMsRdpClientSecuredSettingsDisp = dispinterface
    ['{605BEFCF-39C1-45CC-A811-068FB7BE346D}']
    property KeyboardHookMode: Integer dispid 4;
    property AudioRedirectionMode: Integer dispid 5;
    property StartProgram: WideString dispid 1;
    property WorkDir: WideString dispid 2;
    property FullScreen: Integer dispid 3;
  end;

// *********************************************************************//
// Interface: IMsRdpClientNonScriptable
// Flags:     (0)
// GUID:      {2F079C4C-87B2-4AFD-97AB-20CDB43038AE}
// *********************************************************************//
  IMsRdpClientNonScriptable = interface(IMsTscNonScriptable)
    ['{2F079C4C-87B2-4AFD-97AB-20CDB43038AE}']
    function NotifyRedirectDeviceChange(wParam: UINT_PTR; lParam: LONG_PTR): HResult; stdcall;
    function SendKeys(numKeys: Integer; var pbArrayKeyUp: WordBool; var plKeyData: Integer): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpClient2
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {E7E17DC4-3B71-4BA7-A8E6-281FFADCA28F}
// *********************************************************************//
  IMsRdpClient2 = interface(IMsRdpClient)
    ['{E7E17DC4-3B71-4BA7-A8E6-281FFADCA28F}']
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2; safecall;
    procedure Set_ConnectedStatusText(const pConnectedStatusText: WideString); safecall;
    function Get_ConnectedStatusText: WideString; safecall;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property ConnectedStatusText: WideString read Get_ConnectedStatusText write Set_ConnectedStatusText;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClient2Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {E7E17DC4-3B71-4BA7-A8E6-281FFADCA28F}
// *********************************************************************//
  IMsRdpClient2Disp = dispinterface
    ['{E7E17DC4-3B71-4BA7-A8E6-281FFADCA28F}']
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 readonly dispid 200;
    property ConnectedStatusText: WideString dispid 201;
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClientAdvancedSettings2
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {9AC42117-2B76-4320-AA44-0E616AB8437B}
// *********************************************************************//
  IMsRdpClientAdvancedSettings2 = interface(IMsRdpClientAdvancedSettings)
    ['{9AC42117-2B76-4320-AA44-0E616AB8437B}']
    function Get_CanAutoReconnect: WordBool; safecall;
    procedure Set_EnableAutoReconnect(pfEnableAutoReconnect: WordBool); safecall;
    function Get_EnableAutoReconnect: WordBool; safecall;
    procedure Set_MaxReconnectAttempts(pMaxReconnectAttempts: Integer); safecall;
    function Get_MaxReconnectAttempts: Integer; safecall;
    property CanAutoReconnect: WordBool read Get_CanAutoReconnect;
    property EnableAutoReconnect: WordBool read Get_EnableAutoReconnect write Set_EnableAutoReconnect;
    property MaxReconnectAttempts: Integer read Get_MaxReconnectAttempts write Set_MaxReconnectAttempts;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientAdvancedSettings2Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {9AC42117-2B76-4320-AA44-0E616AB8437B}
// *********************************************************************//
  IMsRdpClientAdvancedSettings2Disp = dispinterface
    ['{9AC42117-2B76-4320-AA44-0E616AB8437B}']
    property CanAutoReconnect: WordBool readonly dispid 205;
    property EnableAutoReconnect: WordBool dispid 206;
    property MaxReconnectAttempts: Integer dispid 207;
    property SmoothScroll: Integer dispid 101;
    property AcceleratorPassthrough: Integer dispid 102;
    property ShadowBitmap: Integer dispid 103;
    property TransportType: Integer dispid 104;
    property SasSequence: Integer dispid 105;
    property EncryptionEnabled: Integer dispid 106;
    property DedicatedTerminal: Integer dispid 107;
    property RDPPort: Integer dispid 108;
    property EnableMouse: Integer dispid 109;
    property DisableCtrlAltDel: Integer dispid 110;
    property EnableWindowsKey: Integer dispid 111;
    property DoubleClickDetect: Integer dispid 112;
    property MaximizeShell: Integer dispid 113;
    property HotKeyFullScreen: Integer dispid 114;
    property HotKeyCtrlEsc: Integer dispid 115;
    property HotKeyAltEsc: Integer dispid 116;
    property HotKeyAltTab: Integer dispid 117;
    property HotKeyAltShiftTab: Integer dispid 118;
    property HotKeyAltSpace: Integer dispid 119;
    property HotKeyCtrlAltDel: Integer dispid 120;
    property orderDrawThreshold: Integer dispid 123;
    property BitmapCacheSize: Integer dispid 124;
    property BitmapVirtualCacheSize: Integer dispid 125;
    property ScaleBitmapCachesByBPP: Integer dispid 175;
    property NumBitmapCaches: Integer dispid 126;
    property CachePersistenceActive: Integer dispid 127;
    property PersistCacheDirectory: WideString writeonly dispid 138;
    property brushSupportLevel: Integer dispid 156;
    property minInputSendInterval: Integer dispid 157;
    property InputEventsAtOnce: Integer dispid 158;
    property maxEventCount: Integer dispid 159;
    property keepAliveInterval: Integer dispid 160;
    property shutdownTimeout: Integer dispid 163;
    property overallConnectionTimeout: Integer dispid 164;
    property singleConnectionTimeout: Integer dispid 165;
    property KeyboardType: Integer dispid 166;
    property KeyboardSubType: Integer dispid 167;
    property KeyboardFunctionKey: Integer dispid 168;
    property WinceFixedPalette: Integer dispid 169;
    property ConnectToServerConsole: WordBool dispid 178;
    property BitmapPersistence: Integer dispid 182;
    property MinutesToIdleTimeout: Integer dispid 183;
    property SmartSizing: WordBool dispid 184;
    property RdpdrLocalPrintingDocName: WideString dispid 185;
    property RdpdrClipCleanTempDirString: WideString dispid 201;
    property RdpdrClipPasteInfoString: WideString dispid 202;
    property ClearTextPassword: WideString writeonly dispid 186;
    property DisplayConnectionBar: WordBool dispid 187;
    property PinConnectionBar: WordBool dispid 188;
    property GrabFocusOnConnect: WordBool dispid 189;
    property LoadBalanceInfo: WideString dispid 190;
    property RedirectDrives: WordBool dispid 191;
    property RedirectPrinters: WordBool dispid 192;
    property RedirectPorts: WordBool dispid 193;
    property RedirectSmartCards: WordBool dispid 194;
    property BitmapVirtualCache16BppSize: Integer dispid 195;
    property BitmapVirtualCache24BppSize: Integer dispid 196;
    property PerformanceFlags: Integer dispid 200;
    property ConnectWithEndpoint: {NOT_OLEAUTO(POleVariant1)}OleVariant writeonly dispid 203;
    property NotifyTSPublicKey: WordBool dispid 204;
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: IMsRdpClient3
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {91B7CBC5-A72E-4FA0-9300-D647D7E897FF}
// *********************************************************************//
  IMsRdpClient3 = interface(IMsRdpClient2)
    ['{91B7CBC5-A72E-4FA0-9300-D647D7E897FF}']
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3; safecall;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClient3Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {91B7CBC5-A72E-4FA0-9300-D647D7E897FF}
// *********************************************************************//
  IMsRdpClient3Disp = dispinterface
    ['{91B7CBC5-A72E-4FA0-9300-D647D7E897FF}']
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 readonly dispid 300;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 readonly dispid 200;
    property ConnectedStatusText: WideString dispid 201;
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClientAdvancedSettings3
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {19CD856B-C542-4C53-ACEE-F127E3BE1A59}
// *********************************************************************//
  IMsRdpClientAdvancedSettings3 = interface(IMsRdpClientAdvancedSettings2)
    ['{19CD856B-C542-4C53-ACEE-F127E3BE1A59}']
    procedure Set_ConnectionBarShowMinimizeButton(pfShowMinimize: WordBool); safecall;
    function Get_ConnectionBarShowMinimizeButton: WordBool; safecall;
    procedure Set_ConnectionBarShowRestoreButton(pfShowRestore: WordBool); safecall;
    function Get_ConnectionBarShowRestoreButton: WordBool; safecall;
    property ConnectionBarShowMinimizeButton: WordBool read Get_ConnectionBarShowMinimizeButton write Set_ConnectionBarShowMinimizeButton;
    property ConnectionBarShowRestoreButton: WordBool read Get_ConnectionBarShowRestoreButton write Set_ConnectionBarShowRestoreButton;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientAdvancedSettings3Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {19CD856B-C542-4C53-ACEE-F127E3BE1A59}
// *********************************************************************//
  IMsRdpClientAdvancedSettings3Disp = dispinterface
    ['{19CD856B-C542-4C53-ACEE-F127E3BE1A59}']
    property ConnectionBarShowMinimizeButton: WordBool dispid 210;
    property ConnectionBarShowRestoreButton: WordBool dispid 211;
    property CanAutoReconnect: WordBool readonly dispid 205;
    property EnableAutoReconnect: WordBool dispid 206;
    property MaxReconnectAttempts: Integer dispid 207;
    property SmoothScroll: Integer dispid 101;
    property AcceleratorPassthrough: Integer dispid 102;
    property ShadowBitmap: Integer dispid 103;
    property TransportType: Integer dispid 104;
    property SasSequence: Integer dispid 105;
    property EncryptionEnabled: Integer dispid 106;
    property DedicatedTerminal: Integer dispid 107;
    property RDPPort: Integer dispid 108;
    property EnableMouse: Integer dispid 109;
    property DisableCtrlAltDel: Integer dispid 110;
    property EnableWindowsKey: Integer dispid 111;
    property DoubleClickDetect: Integer dispid 112;
    property MaximizeShell: Integer dispid 113;
    property HotKeyFullScreen: Integer dispid 114;
    property HotKeyCtrlEsc: Integer dispid 115;
    property HotKeyAltEsc: Integer dispid 116;
    property HotKeyAltTab: Integer dispid 117;
    property HotKeyAltShiftTab: Integer dispid 118;
    property HotKeyAltSpace: Integer dispid 119;
    property HotKeyCtrlAltDel: Integer dispid 120;
    property orderDrawThreshold: Integer dispid 123;
    property BitmapCacheSize: Integer dispid 124;
    property BitmapVirtualCacheSize: Integer dispid 125;
    property ScaleBitmapCachesByBPP: Integer dispid 175;
    property NumBitmapCaches: Integer dispid 126;
    property CachePersistenceActive: Integer dispid 127;
    property PersistCacheDirectory: WideString writeonly dispid 138;
    property brushSupportLevel: Integer dispid 156;
    property minInputSendInterval: Integer dispid 157;
    property InputEventsAtOnce: Integer dispid 158;
    property maxEventCount: Integer dispid 159;
    property keepAliveInterval: Integer dispid 160;
    property shutdownTimeout: Integer dispid 163;
    property overallConnectionTimeout: Integer dispid 164;
    property singleConnectionTimeout: Integer dispid 165;
    property KeyboardType: Integer dispid 166;
    property KeyboardSubType: Integer dispid 167;
    property KeyboardFunctionKey: Integer dispid 168;
    property WinceFixedPalette: Integer dispid 169;
    property ConnectToServerConsole: WordBool dispid 178;
    property BitmapPersistence: Integer dispid 182;
    property MinutesToIdleTimeout: Integer dispid 183;
    property SmartSizing: WordBool dispid 184;
    property RdpdrLocalPrintingDocName: WideString dispid 185;
    property RdpdrClipCleanTempDirString: WideString dispid 201;
    property RdpdrClipPasteInfoString: WideString dispid 202;
    property ClearTextPassword: WideString writeonly dispid 186;
    property DisplayConnectionBar: WordBool dispid 187;
    property PinConnectionBar: WordBool dispid 188;
    property GrabFocusOnConnect: WordBool dispid 189;
    property LoadBalanceInfo: WideString dispid 190;
    property RedirectDrives: WordBool dispid 191;
    property RedirectPrinters: WordBool dispid 192;
    property RedirectPorts: WordBool dispid 193;
    property RedirectSmartCards: WordBool dispid 194;
    property BitmapVirtualCache16BppSize: Integer dispid 195;
    property BitmapVirtualCache24BppSize: Integer dispid 196;
    property PerformanceFlags: Integer dispid 200;
    property ConnectWithEndpoint: {NOT_OLEAUTO(POleVariant1)}OleVariant writeonly dispid 203;
    property NotifyTSPublicKey: WordBool dispid 204;
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: IMsRdpClient4
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {095E0738-D97D-488B-B9F6-DD0E8D66C0DE}
// *********************************************************************//
  IMsRdpClient4 = interface(IMsRdpClient3)
    ['{095E0738-D97D-488B-B9F6-DD0E8D66C0DE}']
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4; safecall;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClient4Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {095E0738-D97D-488B-B9F6-DD0E8D66C0DE}
// *********************************************************************//
  IMsRdpClient4Disp = dispinterface
    ['{095E0738-D97D-488B-B9F6-DD0E8D66C0DE}']
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 readonly dispid 400;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 readonly dispid 300;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 readonly dispid 200;
    property ConnectedStatusText: WideString dispid 201;
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClientNonScriptable2
// Flags:     (0)
// GUID:      {17A5E535-4072-4FA4-AF32-C8D0D47345E9}
// *********************************************************************//
  IMsRdpClientNonScriptable2 = interface(IMsRdpClientNonScriptable)
    ['{17A5E535-4072-4FA4-AF32-C8D0D47345E9}']
    function Set_UIParentWindowHandle(phwndUIParentWindowHandle: wireHWND): HResult; stdcall;
    function Get_UIParentWindowHandle(out phwndUIParentWindowHandle: wireHWND): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpClientAdvancedSettings4
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {FBA7F64E-7345-4405-AE50-FA4A763DC0DE}
// *********************************************************************//
  IMsRdpClientAdvancedSettings4 = interface(IMsRdpClientAdvancedSettings3)
    ['{FBA7F64E-7345-4405-AE50-FA4A763DC0DE}']
    procedure Set_AuthenticationLevel(puiAuthLevel: SYSUINT); safecall;
    function Get_AuthenticationLevel: SYSUINT; safecall;
    property AuthenticationLevel: SYSUINT read Get_AuthenticationLevel write Set_AuthenticationLevel;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientAdvancedSettings4Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {FBA7F64E-7345-4405-AE50-FA4A763DC0DE}
// *********************************************************************//
  IMsRdpClientAdvancedSettings4Disp = dispinterface
    ['{FBA7F64E-7345-4405-AE50-FA4A763DC0DE}']
    property AuthenticationLevel: SYSUINT dispid 212;
    property ConnectionBarShowMinimizeButton: WordBool dispid 210;
    property ConnectionBarShowRestoreButton: WordBool dispid 211;
    property CanAutoReconnect: WordBool readonly dispid 205;
    property EnableAutoReconnect: WordBool dispid 206;
    property MaxReconnectAttempts: Integer dispid 207;
    property SmoothScroll: Integer dispid 101;
    property AcceleratorPassthrough: Integer dispid 102;
    property ShadowBitmap: Integer dispid 103;
    property TransportType: Integer dispid 104;
    property SasSequence: Integer dispid 105;
    property EncryptionEnabled: Integer dispid 106;
    property DedicatedTerminal: Integer dispid 107;
    property RDPPort: Integer dispid 108;
    property EnableMouse: Integer dispid 109;
    property DisableCtrlAltDel: Integer dispid 110;
    property EnableWindowsKey: Integer dispid 111;
    property DoubleClickDetect: Integer dispid 112;
    property MaximizeShell: Integer dispid 113;
    property HotKeyFullScreen: Integer dispid 114;
    property HotKeyCtrlEsc: Integer dispid 115;
    property HotKeyAltEsc: Integer dispid 116;
    property HotKeyAltTab: Integer dispid 117;
    property HotKeyAltShiftTab: Integer dispid 118;
    property HotKeyAltSpace: Integer dispid 119;
    property HotKeyCtrlAltDel: Integer dispid 120;
    property orderDrawThreshold: Integer dispid 123;
    property BitmapCacheSize: Integer dispid 124;
    property BitmapVirtualCacheSize: Integer dispid 125;
    property ScaleBitmapCachesByBPP: Integer dispid 175;
    property NumBitmapCaches: Integer dispid 126;
    property CachePersistenceActive: Integer dispid 127;
    property PersistCacheDirectory: WideString writeonly dispid 138;
    property brushSupportLevel: Integer dispid 156;
    property minInputSendInterval: Integer dispid 157;
    property InputEventsAtOnce: Integer dispid 158;
    property maxEventCount: Integer dispid 159;
    property keepAliveInterval: Integer dispid 160;
    property shutdownTimeout: Integer dispid 163;
    property overallConnectionTimeout: Integer dispid 164;
    property singleConnectionTimeout: Integer dispid 165;
    property KeyboardType: Integer dispid 166;
    property KeyboardSubType: Integer dispid 167;
    property KeyboardFunctionKey: Integer dispid 168;
    property WinceFixedPalette: Integer dispid 169;
    property ConnectToServerConsole: WordBool dispid 178;
    property BitmapPersistence: Integer dispid 182;
    property MinutesToIdleTimeout: Integer dispid 183;
    property SmartSizing: WordBool dispid 184;
    property RdpdrLocalPrintingDocName: WideString dispid 185;
    property RdpdrClipCleanTempDirString: WideString dispid 201;
    property RdpdrClipPasteInfoString: WideString dispid 202;
    property ClearTextPassword: WideString writeonly dispid 186;
    property DisplayConnectionBar: WordBool dispid 187;
    property PinConnectionBar: WordBool dispid 188;
    property GrabFocusOnConnect: WordBool dispid 189;
    property LoadBalanceInfo: WideString dispid 190;
    property RedirectDrives: WordBool dispid 191;
    property RedirectPrinters: WordBool dispid 192;
    property RedirectPorts: WordBool dispid 193;
    property RedirectSmartCards: WordBool dispid 194;
    property BitmapVirtualCache16BppSize: Integer dispid 195;
    property BitmapVirtualCache24BppSize: Integer dispid 196;
    property PerformanceFlags: Integer dispid 200;
    property ConnectWithEndpoint: {NOT_OLEAUTO(POleVariant1)}OleVariant writeonly dispid 203;
    property NotifyTSPublicKey: WordBool dispid 204;
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: IMsRdpClient5
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {4EB5335B-6429-477D-B922-E06A28ECD8BF}
// *********************************************************************//
  IMsRdpClient5 = interface(IMsRdpClient4)
    ['{4EB5335B-6429-477D-B922-E06A28ECD8BF}']
    function Get_TransportSettings: IMsRdpClientTransportSettings; safecall;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5; safecall;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString; safecall;
    function Get_RemoteProgram: ITSRemoteProgram; safecall;
    function Get_MsRdpClientShell: IMsRdpClientShell; safecall;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClient5Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {4EB5335B-6429-477D-B922-E06A28ECD8BF}
// *********************************************************************//
  IMsRdpClient5Disp = dispinterface
    ['{4EB5335B-6429-477D-B922-E06A28ECD8BF}']
    property TransportSettings: IMsRdpClientTransportSettings readonly dispid 500;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 readonly dispid 502;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString; dispid 503;
    property RemoteProgram: ITSRemoteProgram readonly dispid 504;
    property MsRdpClientShell: IMsRdpClientShell readonly dispid 505;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 readonly dispid 400;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 readonly dispid 300;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 readonly dispid 200;
    property ConnectedStatusText: WideString dispid 201;
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClientNonScriptable3
// Flags:     (0)
// GUID:      {B3378D90-0728-45C7-8ED7-B6159FB92219}
// *********************************************************************//
  IMsRdpClientNonScriptable3 = interface(IMsRdpClientNonScriptable2)
    ['{B3378D90-0728-45C7-8ED7-B6159FB92219}']
    function Set_ShowRedirectionWarningDialog(pfShowRdrDlg: WordBool): HResult; stdcall;
    function Get_ShowRedirectionWarningDialog(out pfShowRdrDlg: WordBool): HResult; stdcall;
    function Set_PromptForCredentials(pfPrompt: WordBool): HResult; stdcall;
    function Get_PromptForCredentials(out pfPrompt: WordBool): HResult; stdcall;
    function Set_NegotiateSecurityLayer(pfNegotiate: WordBool): HResult; stdcall;
    function Get_NegotiateSecurityLayer(out pfNegotiate: WordBool): HResult; stdcall;
    function Set_EnableCredSspSupport(pfEnableSupport: WordBool): HResult; stdcall;
    function Get_EnableCredSspSupport(out pfEnableSupport: WordBool): HResult; stdcall;
    function Set_RedirectDynamicDrives(pfRedirectDynamicDrives: WordBool): HResult; stdcall;
    function Get_RedirectDynamicDrives(out pfRedirectDynamicDrives: WordBool): HResult; stdcall;
    function Set_RedirectDynamicDevices(pfRedirectDynamicDevices: WordBool): HResult; stdcall;
    function Get_RedirectDynamicDevices(out pfRedirectDynamicDevices: WordBool): HResult; stdcall;
    function Get_DeviceCollection(out ppDeviceCollection: IMsRdpDeviceCollection): HResult; stdcall;
    function Get_DriveCollection(out ppDeviceCollection: IMsRdpDriveCollection): HResult; stdcall;
    function Set_WarnAboutSendingCredentials(pfWarn: WordBool): HResult; stdcall;
    function Get_WarnAboutSendingCredentials(out pfWarn: WordBool): HResult; stdcall;
    function Set_WarnAboutClipboardRedirection(pfWarn: WordBool): HResult; stdcall;
    function Get_WarnAboutClipboardRedirection(out pfWarn: WordBool): HResult; stdcall;
    function Set_ConnectionBarText(const pConnectionBarText: WideString): HResult; stdcall;
    function Get_ConnectionBarText(out pConnectionBarText: WideString): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpClientTransportSettings
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {720298C0-A099-46F5-9F82-96921BAE4701}
// *********************************************************************//
  IMsRdpClientTransportSettings = interface(IDispatch)
    ['{720298C0-A099-46F5-9F82-96921BAE4701}']
    procedure Set_GatewayHostname(const pProxyHostname: WideString); safecall;
    function Get_GatewayHostname: WideString; safecall;
    procedure Set_GatewayUsageMethod(pulProxyUsageMethod: LongWord); safecall;
    function Get_GatewayUsageMethod: LongWord; safecall;
    procedure Set_GatewayProfileUsageMethod(pulProxyProfileUsageMethod: LongWord); safecall;
    function Get_GatewayProfileUsageMethod: LongWord; safecall;
    procedure Set_GatewayCredsSource(pulProxyCredsSource: LongWord); safecall;
    function Get_GatewayCredsSource: LongWord; safecall;
    procedure Set_GatewayUserSelectedCredsSource(pulProxyCredsSource: LongWord); safecall;
    function Get_GatewayUserSelectedCredsSource: LongWord; safecall;
    function Get_GatewayIsSupported: Integer; safecall;
    function Get_GatewayDefaultUsageMethod: LongWord; safecall;
    property GatewayHostname: WideString read Get_GatewayHostname write Set_GatewayHostname;
    property GatewayUsageMethod: LongWord read Get_GatewayUsageMethod write Set_GatewayUsageMethod;
    property GatewayProfileUsageMethod: LongWord read Get_GatewayProfileUsageMethod write Set_GatewayProfileUsageMethod;
    property GatewayCredsSource: LongWord read Get_GatewayCredsSource write Set_GatewayCredsSource;
    property GatewayUserSelectedCredsSource: LongWord read Get_GatewayUserSelectedCredsSource write Set_GatewayUserSelectedCredsSource;
    property GatewayIsSupported: Integer read Get_GatewayIsSupported;
    property GatewayDefaultUsageMethod: LongWord read Get_GatewayDefaultUsageMethod;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientTransportSettingsDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {720298C0-A099-46F5-9F82-96921BAE4701}
// *********************************************************************//
  IMsRdpClientTransportSettingsDisp = dispinterface
    ['{720298C0-A099-46F5-9F82-96921BAE4701}']
    property GatewayHostname: WideString dispid 210;
    property GatewayUsageMethod: LongWord dispid 211;
    property GatewayProfileUsageMethod: LongWord dispid 212;
    property GatewayCredsSource: LongWord dispid 213;
    property GatewayUserSelectedCredsSource: LongWord dispid 216;
    property GatewayIsSupported: Integer readonly dispid 214;
    property GatewayDefaultUsageMethod: LongWord readonly dispid 215;
  end;

// *********************************************************************//
// Interface: IMsRdpClientAdvancedSettings5
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {FBA7F64E-6783-4405-DA45-FA4A763DABD0}
// *********************************************************************//
  IMsRdpClientAdvancedSettings5 = interface(IMsRdpClientAdvancedSettings4)
    ['{FBA7F64E-6783-4405-DA45-FA4A763DABD0}']
    procedure Set_RedirectClipboard(pfRedirectClipboard: WordBool); safecall;
    function Get_RedirectClipboard: WordBool; safecall;
    procedure Set_AudioRedirectionMode(puiAudioRedirectionMode: SYSUINT); safecall;
    function Get_AudioRedirectionMode: SYSUINT; safecall;
    procedure Set_ConnectionBarShowPinButton(pfShowPin: WordBool); safecall;
    function Get_ConnectionBarShowPinButton: WordBool; safecall;
    procedure Set_PublicMode(pfPublicMode: WordBool); safecall;
    function Get_PublicMode: WordBool; safecall;
    procedure Set_RedirectDevices(pfRedirectPnPDevices: WordBool); safecall;
    function Get_RedirectDevices: WordBool; safecall;
    procedure Set_RedirectPOSDevices(pfRedirectPOSDevices: WordBool); safecall;
    function Get_RedirectPOSDevices: WordBool; safecall;
    procedure Set_BitmapVirtualCache32BppSize(pBitmapVirtualCache32BppSize: Integer); safecall;
    function Get_BitmapVirtualCache32BppSize: Integer; safecall;
    property RedirectClipboard: WordBool read Get_RedirectClipboard write Set_RedirectClipboard;
    property AudioRedirectionMode: SYSUINT read Get_AudioRedirectionMode write Set_AudioRedirectionMode;
    property ConnectionBarShowPinButton: WordBool read Get_ConnectionBarShowPinButton write Set_ConnectionBarShowPinButton;
    property PublicMode: WordBool read Get_PublicMode write Set_PublicMode;
    property RedirectDevices: WordBool read Get_RedirectDevices write Set_RedirectDevices;
    property RedirectPOSDevices: WordBool read Get_RedirectPOSDevices write Set_RedirectPOSDevices;
    property BitmapVirtualCache32BppSize: Integer read Get_BitmapVirtualCache32BppSize write Set_BitmapVirtualCache32BppSize;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientAdvancedSettings5Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {FBA7F64E-6783-4405-DA45-FA4A763DABD0}
// *********************************************************************//
  IMsRdpClientAdvancedSettings5Disp = dispinterface
    ['{FBA7F64E-6783-4405-DA45-FA4A763DABD0}']
    property RedirectClipboard: WordBool dispid 213;
    property AudioRedirectionMode: SYSUINT dispid 215;
    property ConnectionBarShowPinButton: WordBool dispid 216;
    property PublicMode: WordBool dispid 217;
    property RedirectDevices: WordBool dispid 218;
    property RedirectPOSDevices: WordBool dispid 219;
    property BitmapVirtualCache32BppSize: Integer dispid 220;
    property AuthenticationLevel: SYSUINT dispid 212;
    property ConnectionBarShowMinimizeButton: WordBool dispid 210;
    property ConnectionBarShowRestoreButton: WordBool dispid 211;
    property CanAutoReconnect: WordBool readonly dispid 205;
    property EnableAutoReconnect: WordBool dispid 206;
    property MaxReconnectAttempts: Integer dispid 207;
    property SmoothScroll: Integer dispid 101;
    property AcceleratorPassthrough: Integer dispid 102;
    property ShadowBitmap: Integer dispid 103;
    property TransportType: Integer dispid 104;
    property SasSequence: Integer dispid 105;
    property EncryptionEnabled: Integer dispid 106;
    property DedicatedTerminal: Integer dispid 107;
    property RDPPort: Integer dispid 108;
    property EnableMouse: Integer dispid 109;
    property DisableCtrlAltDel: Integer dispid 110;
    property EnableWindowsKey: Integer dispid 111;
    property DoubleClickDetect: Integer dispid 112;
    property MaximizeShell: Integer dispid 113;
    property HotKeyFullScreen: Integer dispid 114;
    property HotKeyCtrlEsc: Integer dispid 115;
    property HotKeyAltEsc: Integer dispid 116;
    property HotKeyAltTab: Integer dispid 117;
    property HotKeyAltShiftTab: Integer dispid 118;
    property HotKeyAltSpace: Integer dispid 119;
    property HotKeyCtrlAltDel: Integer dispid 120;
    property orderDrawThreshold: Integer dispid 123;
    property BitmapCacheSize: Integer dispid 124;
    property BitmapVirtualCacheSize: Integer dispid 125;
    property ScaleBitmapCachesByBPP: Integer dispid 175;
    property NumBitmapCaches: Integer dispid 126;
    property CachePersistenceActive: Integer dispid 127;
    property PersistCacheDirectory: WideString writeonly dispid 138;
    property brushSupportLevel: Integer dispid 156;
    property minInputSendInterval: Integer dispid 157;
    property InputEventsAtOnce: Integer dispid 158;
    property maxEventCount: Integer dispid 159;
    property keepAliveInterval: Integer dispid 160;
    property shutdownTimeout: Integer dispid 163;
    property overallConnectionTimeout: Integer dispid 164;
    property singleConnectionTimeout: Integer dispid 165;
    property KeyboardType: Integer dispid 166;
    property KeyboardSubType: Integer dispid 167;
    property KeyboardFunctionKey: Integer dispid 168;
    property WinceFixedPalette: Integer dispid 169;
    property ConnectToServerConsole: WordBool dispid 178;
    property BitmapPersistence: Integer dispid 182;
    property MinutesToIdleTimeout: Integer dispid 183;
    property SmartSizing: WordBool dispid 184;
    property RdpdrLocalPrintingDocName: WideString dispid 185;
    property RdpdrClipCleanTempDirString: WideString dispid 201;
    property RdpdrClipPasteInfoString: WideString dispid 202;
    property ClearTextPassword: WideString writeonly dispid 186;
    property DisplayConnectionBar: WordBool dispid 187;
    property PinConnectionBar: WordBool dispid 188;
    property GrabFocusOnConnect: WordBool dispid 189;
    property LoadBalanceInfo: WideString dispid 190;
    property RedirectDrives: WordBool dispid 191;
    property RedirectPrinters: WordBool dispid 192;
    property RedirectPorts: WordBool dispid 193;
    property RedirectSmartCards: WordBool dispid 194;
    property BitmapVirtualCache16BppSize: Integer dispid 195;
    property BitmapVirtualCache24BppSize: Integer dispid 196;
    property PerformanceFlags: Integer dispid 200;
    property ConnectWithEndpoint: {NOT_OLEAUTO(POleVariant1)}OleVariant writeonly dispid 203;
    property NotifyTSPublicKey: WordBool dispid 204;
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: ITSRemoteProgram
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {FDD029F9-467A-4C49-8529-64B521DBD1B4}
// *********************************************************************//
  ITSRemoteProgram = interface(IDispatch)
    ['{FDD029F9-467A-4C49-8529-64B521DBD1B4}']
    procedure Set_RemoteProgramMode(pvboolRemoteProgramMode: WordBool); safecall;
    function Get_RemoteProgramMode: WordBool; safecall;
    procedure ServerStartProgram(const bstrExecutablePath: WideString; 
                                 const bstrFilePath: WideString; 
                                 const bstrWorkingDirectory: WideString; 
                                 vbExpandEnvVarInWorkingDirectoryOnServer: WordBool; 
                                 const bstrArguments: WideString; 
                                 vbExpandEnvVarInArgumentsOnServer: WordBool); safecall;
    property RemoteProgramMode: WordBool read Get_RemoteProgramMode write Set_RemoteProgramMode;
  end;

// *********************************************************************//
// DispIntf:  ITSRemoteProgramDisp
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {FDD029F9-467A-4C49-8529-64B521DBD1B4}
// *********************************************************************//
  ITSRemoteProgramDisp = dispinterface
    ['{FDD029F9-467A-4C49-8529-64B521DBD1B4}']
    property RemoteProgramMode: WordBool dispid 200;
    procedure ServerStartProgram(const bstrExecutablePath: WideString; 
                                 const bstrFilePath: WideString; 
                                 const bstrWorkingDirectory: WideString; 
                                 vbExpandEnvVarInWorkingDirectoryOnServer: WordBool; 
                                 const bstrArguments: WideString; 
                                 vbExpandEnvVarInArgumentsOnServer: WordBool); dispid 201;
  end;

// *********************************************************************//
// Interface: IMsRdpClientShell
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {D012AE6D-C19A-4BFE-B367-201F8911F134}
// *********************************************************************//
  IMsRdpClientShell = interface(IDispatch)
    ['{D012AE6D-C19A-4BFE-B367-201F8911F134}']
    procedure Launch; safecall;
    procedure Set_RdpFileContents(const pszRdpFile: WideString); safecall;
    function Get_RdpFileContents: WideString; safecall;
    procedure SetRdpProperty(const szProperty: WideString; Value: OleVariant); safecall;
    function GetRdpProperty(const szProperty: WideString): OleVariant; safecall;
    function Get_IsRemoteProgramClientInstalled: WordBool; safecall;
    procedure Set_PublicMode(pfPublicMode: WordBool); safecall;
    function Get_PublicMode: WordBool; safecall;
    procedure ShowTrustedSitesManagementDialog; safecall;
    property RdpFileContents: WideString read Get_RdpFileContents write Set_RdpFileContents;
    property IsRemoteProgramClientInstalled: WordBool read Get_IsRemoteProgramClientInstalled;
    property PublicMode: WordBool read Get_PublicMode write Set_PublicMode;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientShellDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {D012AE6D-C19A-4BFE-B367-201F8911F134}
// *********************************************************************//
  IMsRdpClientShellDisp = dispinterface
    ['{D012AE6D-C19A-4BFE-B367-201F8911F134}']
    procedure Launch; dispid 201;
    property RdpFileContents: WideString dispid 202;
    procedure SetRdpProperty(const szProperty: WideString; Value: OleVariant); dispid 203;
    function GetRdpProperty(const szProperty: WideString): OleVariant; dispid 204;
    property IsRemoteProgramClientInstalled: WordBool readonly dispid 205;
    property PublicMode: WordBool dispid 211;
    procedure ShowTrustedSitesManagementDialog; dispid 212;
  end;

// *********************************************************************//
// Interface: IMsRdpDeviceCollection
// Flags:     (0)
// GUID:      {56540617-D281-488C-8738-6A8FDF64A118}
// *********************************************************************//
  IMsRdpDeviceCollection = interface(IUnknown)
    ['{56540617-D281-488C-8738-6A8FDF64A118}']
    function RescanDevices(vboolDynRedir: WordBool): HResult; stdcall;
    function Get_DeviceByIndex(index: LongWord; out ppDevice: IMsRdpDevice): HResult; stdcall;
    function Get_DeviceById(const devInstanceId: WideString; out ppDevice: IMsRdpDevice): HResult; stdcall;
    function Get_DeviceCount(out pDeviceCount: LongWord): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpDevice
// Flags:     (0)
// GUID:      {60C3B9C8-9E92-4F5E-A3E7-604A912093EA}
// *********************************************************************//
  IMsRdpDevice = interface(IUnknown)
    ['{60C3B9C8-9E92-4F5E-A3E7-604A912093EA}']
    function Get_DeviceInstanceId(out pDevInstanceId: WideString): HResult; stdcall;
    function Get_FriendlyName(out pFriendlyName: WideString): HResult; stdcall;
    function Get_DeviceDescription(out pDeviceDescription: WideString): HResult; stdcall;
    function Set_RedirectionState(pvboolRedirState: WordBool): HResult; stdcall;
    function Get_RedirectionState(out pvboolRedirState: WordBool): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpDriveCollection
// Flags:     (0)
// GUID:      {7FF17599-DA2C-4677-AD35-F60C04FE1585}
// *********************************************************************//
  IMsRdpDriveCollection = interface(IUnknown)
    ['{7FF17599-DA2C-4677-AD35-F60C04FE1585}']
    function RescanDrives(vboolDynRedir: WordBool): HResult; stdcall;
    function Get_DriveByIndex(index: LongWord; out ppDevice: IMsRdpDrive): HResult; stdcall;
    function Get_DriveCount(out pDriveCount: LongWord): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpDrive
// Flags:     (0)
// GUID:      {D28B5458-F694-47A8-8E61-40356A767E46}
// *********************************************************************//
  IMsRdpDrive = interface(IUnknown)
    ['{D28B5458-F694-47A8-8E61-40356A767E46}']
    function Get_Name(out pName: WideString): HResult; stdcall;
    function Set_RedirectionState(pvboolRedirState: WordBool): HResult; stdcall;
    function Get_RedirectionState(out pvboolRedirState: WordBool): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpClient6
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {D43B7D80-8517-4B6D-9EAC-96AD6800D7F2}
// *********************************************************************//
  IMsRdpClient6 = interface(IMsRdpClient5)
    ['{D43B7D80-8517-4B6D-9EAC-96AD6800D7F2}']
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6; safecall;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2; safecall;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClient6Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {D43B7D80-8517-4B6D-9EAC-96AD6800D7F2}
// *********************************************************************//
  IMsRdpClient6Disp = dispinterface
    ['{D43B7D80-8517-4B6D-9EAC-96AD6800D7F2}']
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 readonly dispid 507;
    property TransportSettings2: IMsRdpClientTransportSettings2 readonly dispid 506;
    property TransportSettings: IMsRdpClientTransportSettings readonly dispid 500;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 readonly dispid 502;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString; dispid 503;
    property RemoteProgram: ITSRemoteProgram readonly dispid 504;
    property MsRdpClientShell: IMsRdpClientShell readonly dispid 505;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 readonly dispid 400;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 readonly dispid 300;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 readonly dispid 200;
    property ConnectedStatusText: WideString dispid 201;
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClientNonScriptable4
// Flags:     (0)
// GUID:      {F50FA8AA-1C7D-4F59-B15C-A90CACAE1FCB}
// *********************************************************************//
  IMsRdpClientNonScriptable4 = interface(IMsRdpClientNonScriptable3)
    ['{F50FA8AA-1C7D-4F59-B15C-A90CACAE1FCB}']
    function Set_RedirectionWarningType(pWrnType: RedirectionWarningType): HResult; stdcall;
    function Get_RedirectionWarningType(out pWrnType: RedirectionWarningType): HResult; stdcall;
    function Set_MarkRdpSettingsSecure(pfRdpSecure: WordBool): HResult; stdcall;
    function Get_MarkRdpSettingsSecure(out pfRdpSecure: WordBool): HResult; stdcall;
    function Set_PublisherCertificateChain(const pVarCert: OleVariant): HResult; stdcall;
    function Get_PublisherCertificateChain(out pVarCert: OleVariant): HResult; stdcall;
    function Set_WarnAboutPrinterRedirection(pfWarn: WordBool): HResult; stdcall;
    function Get_WarnAboutPrinterRedirection(out pfWarn: WordBool): HResult; stdcall;
    function Set_AllowCredentialSaving(pfAllowSave: WordBool): HResult; stdcall;
    function Get_AllowCredentialSaving(out pfAllowSave: WordBool): HResult; stdcall;
    function Set_PromptForCredsOnClient(pfPromptForCredsOnClient: WordBool): HResult; stdcall;
    function Get_PromptForCredsOnClient(out pfPromptForCredsOnClient: WordBool): HResult; stdcall;
    function Set_LaunchedViaClientShellInterface(pfLaunchedViaClientShellInterface: WordBool): HResult; stdcall;
    function Get_LaunchedViaClientShellInterface(out pfLaunchedViaClientShellInterface: WordBool): HResult; stdcall;
    function Set_TrustedZoneSite(pfIsTrustedZone: WordBool): HResult; stdcall;
    function Get_TrustedZoneSite(out pfIsTrustedZone: WordBool): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpClientAdvancedSettings6
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {222C4B5D-45D9-4DF0-A7C6-60CF9089D285}
// *********************************************************************//
  IMsRdpClientAdvancedSettings6 = interface(IMsRdpClientAdvancedSettings5)
    ['{222C4B5D-45D9-4DF0-A7C6-60CF9089D285}']
    procedure Set_RelativeMouseMode(pfRelativeMouseMode: WordBool); safecall;
    function Get_RelativeMouseMode: WordBool; safecall;
    function Get_AuthenticationServiceClass: WideString; safecall;
    procedure Set_AuthenticationServiceClass(const pbstrAuthServiceClass: WideString); safecall;
    function Get_PCB: WideString; safecall;
    procedure Set_PCB(const bstrPCB: WideString); safecall;
    procedure Set_HotKeyFocusReleaseLeft(HotKeyFocusReleaseLeft: Integer); safecall;
    function Get_HotKeyFocusReleaseLeft: Integer; safecall;
    procedure Set_HotKeyFocusReleaseRight(HotKeyFocusReleaseRight: Integer); safecall;
    function Get_HotKeyFocusReleaseRight: Integer; safecall;
    procedure Set_EnableCredSspSupport(pfEnableSupport: WordBool); safecall;
    function Get_EnableCredSspSupport: WordBool; safecall;
    function Get_AuthenticationType: SYSUINT; safecall;
    procedure Set_ConnectToAdministerServer(pConnectToAdministerServer: WordBool); safecall;
    function Get_ConnectToAdministerServer: WordBool; safecall;
    property RelativeMouseMode: WordBool read Get_RelativeMouseMode write Set_RelativeMouseMode;
    property AuthenticationServiceClass: WideString read Get_AuthenticationServiceClass write Set_AuthenticationServiceClass;
    property PCB: WideString read Get_PCB write Set_PCB;
    property HotKeyFocusReleaseLeft: Integer read Get_HotKeyFocusReleaseLeft write Set_HotKeyFocusReleaseLeft;
    property HotKeyFocusReleaseRight: Integer read Get_HotKeyFocusReleaseRight write Set_HotKeyFocusReleaseRight;
    property EnableCredSspSupport: WordBool read Get_EnableCredSspSupport write Set_EnableCredSspSupport;
    property AuthenticationType: SYSUINT read Get_AuthenticationType;
    property ConnectToAdministerServer: WordBool read Get_ConnectToAdministerServer write Set_ConnectToAdministerServer;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientAdvancedSettings6Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {222C4B5D-45D9-4DF0-A7C6-60CF9089D285}
// *********************************************************************//
  IMsRdpClientAdvancedSettings6Disp = dispinterface
    ['{222C4B5D-45D9-4DF0-A7C6-60CF9089D285}']
    property RelativeMouseMode: WordBool dispid 221;
    property AuthenticationServiceClass: WideString dispid 222;
    property PCB: WideString dispid 223;
    property HotKeyFocusReleaseLeft: Integer dispid 224;
    property HotKeyFocusReleaseRight: Integer dispid 225;
    property EnableCredSspSupport: WordBool dispid 17;
    property AuthenticationType: SYSUINT readonly dispid 226;
    property ConnectToAdministerServer: WordBool dispid 227;
    property RedirectClipboard: WordBool dispid 213;
    property AudioRedirectionMode: SYSUINT dispid 215;
    property ConnectionBarShowPinButton: WordBool dispid 216;
    property PublicMode: WordBool dispid 217;
    property RedirectDevices: WordBool dispid 218;
    property RedirectPOSDevices: WordBool dispid 219;
    property BitmapVirtualCache32BppSize: Integer dispid 220;
    property AuthenticationLevel: SYSUINT dispid 212;
    property ConnectionBarShowMinimizeButton: WordBool dispid 210;
    property ConnectionBarShowRestoreButton: WordBool dispid 211;
    property CanAutoReconnect: WordBool readonly dispid 205;
    property EnableAutoReconnect: WordBool dispid 206;
    property MaxReconnectAttempts: Integer dispid 207;
    property SmoothScroll: Integer dispid 101;
    property AcceleratorPassthrough: Integer dispid 102;
    property ShadowBitmap: Integer dispid 103;
    property TransportType: Integer dispid 104;
    property SasSequence: Integer dispid 105;
    property EncryptionEnabled: Integer dispid 106;
    property DedicatedTerminal: Integer dispid 107;
    property RDPPort: Integer dispid 108;
    property EnableMouse: Integer dispid 109;
    property DisableCtrlAltDel: Integer dispid 110;
    property EnableWindowsKey: Integer dispid 111;
    property DoubleClickDetect: Integer dispid 112;
    property MaximizeShell: Integer dispid 113;
    property HotKeyFullScreen: Integer dispid 114;
    property HotKeyCtrlEsc: Integer dispid 115;
    property HotKeyAltEsc: Integer dispid 116;
    property HotKeyAltTab: Integer dispid 117;
    property HotKeyAltShiftTab: Integer dispid 118;
    property HotKeyAltSpace: Integer dispid 119;
    property HotKeyCtrlAltDel: Integer dispid 120;
    property orderDrawThreshold: Integer dispid 123;
    property BitmapCacheSize: Integer dispid 124;
    property BitmapVirtualCacheSize: Integer dispid 125;
    property ScaleBitmapCachesByBPP: Integer dispid 175;
    property NumBitmapCaches: Integer dispid 126;
    property CachePersistenceActive: Integer dispid 127;
    property PersistCacheDirectory: WideString writeonly dispid 138;
    property brushSupportLevel: Integer dispid 156;
    property minInputSendInterval: Integer dispid 157;
    property InputEventsAtOnce: Integer dispid 158;
    property maxEventCount: Integer dispid 159;
    property keepAliveInterval: Integer dispid 160;
    property shutdownTimeout: Integer dispid 163;
    property overallConnectionTimeout: Integer dispid 164;
    property singleConnectionTimeout: Integer dispid 165;
    property KeyboardType: Integer dispid 166;
    property KeyboardSubType: Integer dispid 167;
    property KeyboardFunctionKey: Integer dispid 168;
    property WinceFixedPalette: Integer dispid 169;
    property ConnectToServerConsole: WordBool dispid 178;
    property BitmapPersistence: Integer dispid 182;
    property MinutesToIdleTimeout: Integer dispid 183;
    property SmartSizing: WordBool dispid 184;
    property RdpdrLocalPrintingDocName: WideString dispid 185;
    property RdpdrClipCleanTempDirString: WideString dispid 201;
    property RdpdrClipPasteInfoString: WideString dispid 202;
    property ClearTextPassword: WideString writeonly dispid 186;
    property DisplayConnectionBar: WordBool dispid 187;
    property PinConnectionBar: WordBool dispid 188;
    property GrabFocusOnConnect: WordBool dispid 189;
    property LoadBalanceInfo: WideString dispid 190;
    property RedirectDrives: WordBool dispid 191;
    property RedirectPrinters: WordBool dispid 192;
    property RedirectPorts: WordBool dispid 193;
    property RedirectSmartCards: WordBool dispid 194;
    property BitmapVirtualCache16BppSize: Integer dispid 195;
    property BitmapVirtualCache24BppSize: Integer dispid 196;
    property PerformanceFlags: Integer dispid 200;
    property ConnectWithEndpoint: {NOT_OLEAUTO(POleVariant1)}OleVariant writeonly dispid 203;
    property NotifyTSPublicKey: WordBool dispid 204;
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: IMsRdpClientTransportSettings2
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {67341688-D606-4C73-A5D2-2E0489009319}
// *********************************************************************//
  IMsRdpClientTransportSettings2 = interface(IMsRdpClientTransportSettings)
    ['{67341688-D606-4C73-A5D2-2E0489009319}']
    procedure Set_GatewayCredSharing(pulProxyCredSharing: LongWord); safecall;
    function Get_GatewayCredSharing: LongWord; safecall;
    procedure Set_GatewayPreAuthRequirement(pulProxyPreAuthRequirement: LongWord); safecall;
    function Get_GatewayPreAuthRequirement: LongWord; safecall;
    procedure Set_GatewayPreAuthServerAddr(const pbstrProxyPreAuthServerAddr: WideString); safecall;
    function Get_GatewayPreAuthServerAddr: WideString; safecall;
    procedure Set_GatewaySupportUrl(const pbstrProxySupportUrl: WideString); safecall;
    function Get_GatewaySupportUrl: WideString; safecall;
    procedure Set_GatewayEncryptedOtpCookie(const pbstrEncryptedOtpCookie: WideString); safecall;
    function Get_GatewayEncryptedOtpCookie: WideString; safecall;
    procedure Set_GatewayEncryptedOtpCookieSize(pulEncryptedOtpCookieSize: LongWord); safecall;
    function Get_GatewayEncryptedOtpCookieSize: LongWord; safecall;
    procedure Set_GatewayUsername(const pProxyUsername: WideString); safecall;
    function Get_GatewayUsername: WideString; safecall;
    procedure Set_GatewayDomain(const pProxyDomain: WideString); safecall;
    function Get_GatewayDomain: WideString; safecall;
    procedure Set_GatewayPassword(const Param1: WideString); safecall;
    property GatewayCredSharing: LongWord read Get_GatewayCredSharing write Set_GatewayCredSharing;
    property GatewayPreAuthRequirement: LongWord read Get_GatewayPreAuthRequirement write Set_GatewayPreAuthRequirement;
    property GatewayPreAuthServerAddr: WideString read Get_GatewayPreAuthServerAddr write Set_GatewayPreAuthServerAddr;
    property GatewaySupportUrl: WideString read Get_GatewaySupportUrl write Set_GatewaySupportUrl;
    property GatewayEncryptedOtpCookie: WideString read Get_GatewayEncryptedOtpCookie write Set_GatewayEncryptedOtpCookie;
    property GatewayEncryptedOtpCookieSize: LongWord read Get_GatewayEncryptedOtpCookieSize write Set_GatewayEncryptedOtpCookieSize;
    property GatewayUsername: WideString read Get_GatewayUsername write Set_GatewayUsername;
    property GatewayDomain: WideString read Get_GatewayDomain write Set_GatewayDomain;
    property GatewayPassword: WideString write Set_GatewayPassword;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientTransportSettings2Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {67341688-D606-4C73-A5D2-2E0489009319}
// *********************************************************************//
  IMsRdpClientTransportSettings2Disp = dispinterface
    ['{67341688-D606-4C73-A5D2-2E0489009319}']
    property GatewayCredSharing: LongWord dispid 222;
    property GatewayPreAuthRequirement: LongWord dispid 217;
    property GatewayPreAuthServerAddr: WideString dispid 218;
    property GatewaySupportUrl: WideString dispid 219;
    property GatewayEncryptedOtpCookie: WideString dispid 220;
    property GatewayEncryptedOtpCookieSize: LongWord dispid 221;
    property GatewayUsername: WideString dispid 223;
    property GatewayDomain: WideString dispid 224;
    property GatewayPassword: WideString writeonly dispid 225;
    property GatewayHostname: WideString dispid 210;
    property GatewayUsageMethod: LongWord dispid 211;
    property GatewayProfileUsageMethod: LongWord dispid 212;
    property GatewayCredsSource: LongWord dispid 213;
    property GatewayUserSelectedCredsSource: LongWord dispid 216;
    property GatewayIsSupported: Integer readonly dispid 214;
    property GatewayDefaultUsageMethod: LongWord readonly dispid 215;
  end;

// *********************************************************************//
// Interface: IMsRdpClient7
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {B2A5B5CE-3461-444A-91D4-ADD26D070638}
// *********************************************************************//
  IMsRdpClient7 = interface(IMsRdpClient6)
    ['{B2A5B5CE-3461-444A-91D4-ADD26D070638}']
    function Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7; safecall;
    function Get_TransportSettings3: IMsRdpClientTransportSettings3; safecall;
    function GetStatusText(statusCode: SYSUINT): WideString; safecall;
    function Get_SecuredSettings3: IMsRdpClientSecuredSettings2; safecall;
    function Get_RemoteProgram2: ITSRemoteProgram2; safecall;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 read Get_AdvancedSettings8;
    property TransportSettings3: IMsRdpClientTransportSettings3 read Get_TransportSettings3;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 read Get_SecuredSettings3;
    property RemoteProgram2: ITSRemoteProgram2 read Get_RemoteProgram2;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClient7Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {B2A5B5CE-3461-444A-91D4-ADD26D070638}
// *********************************************************************//
  IMsRdpClient7Disp = dispinterface
    ['{B2A5B5CE-3461-444A-91D4-ADD26D070638}']
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 readonly dispid 600;
    property TransportSettings3: IMsRdpClientTransportSettings3 readonly dispid 601;
    function GetStatusText(statusCode: SYSUINT): WideString; dispid 602;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 readonly dispid 603;
    property RemoteProgram2: ITSRemoteProgram2 readonly dispid 604;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 readonly dispid 507;
    property TransportSettings2: IMsRdpClientTransportSettings2 readonly dispid 506;
    property TransportSettings: IMsRdpClientTransportSettings readonly dispid 500;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 readonly dispid 502;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString; dispid 503;
    property RemoteProgram: ITSRemoteProgram readonly dispid 504;
    property MsRdpClientShell: IMsRdpClientShell readonly dispid 505;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 readonly dispid 400;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 readonly dispid 300;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 readonly dispid 200;
    property ConnectedStatusText: WideString dispid 201;
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClientNonScriptable5
// Flags:     (0)
// GUID:      {4F6996D5-D7B1-412C-B0FF-063718566907}
// *********************************************************************//
  IMsRdpClientNonScriptable5 = interface(IMsRdpClientNonScriptable4)
    ['{4F6996D5-D7B1-412C-B0FF-063718566907}']
    function Set_UseMultimon(pfUseMultimon: WordBool): HResult; stdcall;
    function Get_UseMultimon(out pfUseMultimon: WordBool): HResult; stdcall;
    function Get_RemoteMonitorCount(out pcRemoteMonitors: LongWord): HResult; stdcall;
    function GetRemoteMonitorsBoundingBox(out pLeft: Integer; out pTop: Integer; 
                                          out pRight: Integer; out pBottom: Integer): HResult; stdcall;
    function Get_RemoteMonitorLayoutMatchesLocal(out pfRemoteMatchesLocal: WordBool): HResult; stdcall;
    function Set_DisableConnectionBar(Param1: WordBool): HResult; stdcall;
    function Set_DisableRemoteAppCapsCheck(pfDisableRemoteAppCapsCheck: WordBool): HResult; stdcall;
    function Get_DisableRemoteAppCapsCheck(out pfDisableRemoteAppCapsCheck: WordBool): HResult; stdcall;
    function Set_WarnAboutDirectXRedirection(pfWarn: WordBool): HResult; stdcall;
    function Get_WarnAboutDirectXRedirection(out pfWarn: WordBool): HResult; stdcall;
    function Set_AllowPromptingForCredentials(pfAllow: WordBool): HResult; stdcall;
    function Get_AllowPromptingForCredentials(out pfAllow: WordBool): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpClientAdvancedSettings7
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {26036036-4010-4578-8091-0DB9A1EDF9C3}
// *********************************************************************//
  IMsRdpClientAdvancedSettings7 = interface(IMsRdpClientAdvancedSettings6)
    ['{26036036-4010-4578-8091-0DB9A1EDF9C3}']
    procedure Set_AudioCaptureRedirectionMode(pfRedir: WordBool); safecall;
    function Get_AudioCaptureRedirectionMode: WordBool; safecall;
    procedure Set_VideoPlaybackMode(pVideoPlaybackMode: SYSUINT); safecall;
    function Get_VideoPlaybackMode: SYSUINT; safecall;
    procedure Set_EnableSuperPan(pfEnableSuperPan: WordBool); safecall;
    function Get_EnableSuperPan: WordBool; safecall;
    procedure Set_SuperPanAccelerationFactor(puAccelFactor: LongWord); safecall;
    function Get_SuperPanAccelerationFactor: LongWord; safecall;
    procedure Set_NegotiateSecurityLayer(pfNegotiate: WordBool); safecall;
    function Get_NegotiateSecurityLayer: WordBool; safecall;
    procedure Set_AudioQualityMode(pAudioQualityMode: SYSUINT); safecall;
    function Get_AudioQualityMode: SYSUINT; safecall;
    procedure Set_RedirectDirectX(pfRedirectDirectX: WordBool); safecall;
    function Get_RedirectDirectX: WordBool; safecall;
    procedure Set_NetworkConnectionType(pConnectionType: SYSUINT); safecall;
    function Get_NetworkConnectionType: SYSUINT; safecall;
    property AudioCaptureRedirectionMode: WordBool read Get_AudioCaptureRedirectionMode write Set_AudioCaptureRedirectionMode;
    property VideoPlaybackMode: SYSUINT read Get_VideoPlaybackMode write Set_VideoPlaybackMode;
    property EnableSuperPan: WordBool read Get_EnableSuperPan write Set_EnableSuperPan;
    property SuperPanAccelerationFactor: LongWord read Get_SuperPanAccelerationFactor write Set_SuperPanAccelerationFactor;
    property NegotiateSecurityLayer: WordBool read Get_NegotiateSecurityLayer write Set_NegotiateSecurityLayer;
    property AudioQualityMode: SYSUINT read Get_AudioQualityMode write Set_AudioQualityMode;
    property RedirectDirectX: WordBool read Get_RedirectDirectX write Set_RedirectDirectX;
    property NetworkConnectionType: SYSUINT read Get_NetworkConnectionType write Set_NetworkConnectionType;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientAdvancedSettings7Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {26036036-4010-4578-8091-0DB9A1EDF9C3}
// *********************************************************************//
  IMsRdpClientAdvancedSettings7Disp = dispinterface
    ['{26036036-4010-4578-8091-0DB9A1EDF9C3}']
    property AudioCaptureRedirectionMode: WordBool dispid 228;
    property VideoPlaybackMode: SYSUINT dispid 229;
    property EnableSuperPan: WordBool dispid 230;
    property SuperPanAccelerationFactor: LongWord dispid 231;
    property NegotiateSecurityLayer: WordBool dispid 232;
    property AudioQualityMode: SYSUINT dispid 233;
    property RedirectDirectX: WordBool dispid 234;
    property NetworkConnectionType: SYSUINT dispid 235;
    property RelativeMouseMode: WordBool dispid 221;
    property AuthenticationServiceClass: WideString dispid 222;
    property PCB: WideString dispid 223;
    property HotKeyFocusReleaseLeft: Integer dispid 224;
    property HotKeyFocusReleaseRight: Integer dispid 225;
    property EnableCredSspSupport: WordBool dispid 17;
    property AuthenticationType: SYSUINT readonly dispid 226;
    property ConnectToAdministerServer: WordBool dispid 227;
    property RedirectClipboard: WordBool dispid 213;
    property AudioRedirectionMode: SYSUINT dispid 215;
    property ConnectionBarShowPinButton: WordBool dispid 216;
    property PublicMode: WordBool dispid 217;
    property RedirectDevices: WordBool dispid 218;
    property RedirectPOSDevices: WordBool dispid 219;
    property BitmapVirtualCache32BppSize: Integer dispid 220;
    property AuthenticationLevel: SYSUINT dispid 212;
    property ConnectionBarShowMinimizeButton: WordBool dispid 210;
    property ConnectionBarShowRestoreButton: WordBool dispid 211;
    property CanAutoReconnect: WordBool readonly dispid 205;
    property EnableAutoReconnect: WordBool dispid 206;
    property MaxReconnectAttempts: Integer dispid 207;
    property SmoothScroll: Integer dispid 101;
    property AcceleratorPassthrough: Integer dispid 102;
    property ShadowBitmap: Integer dispid 103;
    property TransportType: Integer dispid 104;
    property SasSequence: Integer dispid 105;
    property EncryptionEnabled: Integer dispid 106;
    property DedicatedTerminal: Integer dispid 107;
    property RDPPort: Integer dispid 108;
    property EnableMouse: Integer dispid 109;
    property DisableCtrlAltDel: Integer dispid 110;
    property EnableWindowsKey: Integer dispid 111;
    property DoubleClickDetect: Integer dispid 112;
    property MaximizeShell: Integer dispid 113;
    property HotKeyFullScreen: Integer dispid 114;
    property HotKeyCtrlEsc: Integer dispid 115;
    property HotKeyAltEsc: Integer dispid 116;
    property HotKeyAltTab: Integer dispid 117;
    property HotKeyAltShiftTab: Integer dispid 118;
    property HotKeyAltSpace: Integer dispid 119;
    property HotKeyCtrlAltDel: Integer dispid 120;
    property orderDrawThreshold: Integer dispid 123;
    property BitmapCacheSize: Integer dispid 124;
    property BitmapVirtualCacheSize: Integer dispid 125;
    property ScaleBitmapCachesByBPP: Integer dispid 175;
    property NumBitmapCaches: Integer dispid 126;
    property CachePersistenceActive: Integer dispid 127;
    property PersistCacheDirectory: WideString writeonly dispid 138;
    property brushSupportLevel: Integer dispid 156;
    property minInputSendInterval: Integer dispid 157;
    property InputEventsAtOnce: Integer dispid 158;
    property maxEventCount: Integer dispid 159;
    property keepAliveInterval: Integer dispid 160;
    property shutdownTimeout: Integer dispid 163;
    property overallConnectionTimeout: Integer dispid 164;
    property singleConnectionTimeout: Integer dispid 165;
    property KeyboardType: Integer dispid 166;
    property KeyboardSubType: Integer dispid 167;
    property KeyboardFunctionKey: Integer dispid 168;
    property WinceFixedPalette: Integer dispid 169;
    property ConnectToServerConsole: WordBool dispid 178;
    property BitmapPersistence: Integer dispid 182;
    property MinutesToIdleTimeout: Integer dispid 183;
    property SmartSizing: WordBool dispid 184;
    property RdpdrLocalPrintingDocName: WideString dispid 185;
    property RdpdrClipCleanTempDirString: WideString dispid 201;
    property RdpdrClipPasteInfoString: WideString dispid 202;
    property ClearTextPassword: WideString writeonly dispid 186;
    property DisplayConnectionBar: WordBool dispid 187;
    property PinConnectionBar: WordBool dispid 188;
    property GrabFocusOnConnect: WordBool dispid 189;
    property LoadBalanceInfo: WideString dispid 190;
    property RedirectDrives: WordBool dispid 191;
    property RedirectPrinters: WordBool dispid 192;
    property RedirectPorts: WordBool dispid 193;
    property RedirectSmartCards: WordBool dispid 194;
    property BitmapVirtualCache16BppSize: Integer dispid 195;
    property BitmapVirtualCache24BppSize: Integer dispid 196;
    property PerformanceFlags: Integer dispid 200;
    property ConnectWithEndpoint: {NOT_OLEAUTO(POleVariant1)}OleVariant writeonly dispid 203;
    property NotifyTSPublicKey: WordBool dispid 204;
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: IMsRdpClientTransportSettings3
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {3D5B21AC-748D-41DE-8F30-E15169586BD4}
// *********************************************************************//
  IMsRdpClientTransportSettings3 = interface(IMsRdpClientTransportSettings2)
    ['{3D5B21AC-748D-41DE-8F30-E15169586BD4}']
    procedure Set_GatewayCredSourceCookie(pulProxyCredSourceCookie: LongWord); safecall;
    function Get_GatewayCredSourceCookie: LongWord; safecall;
    procedure Set_GatewayAuthCookieServerAddr(const pbstrProxyAuthCookieServerAddr: WideString); safecall;
    function Get_GatewayAuthCookieServerAddr: WideString; safecall;
    procedure Set_GatewayEncryptedAuthCookie(const pbstrEncryptedAuthCookie: WideString); safecall;
    function Get_GatewayEncryptedAuthCookie: WideString; safecall;
    procedure Set_GatewayEncryptedAuthCookieSize(pulEncryptedAuthCookieSize: LongWord); safecall;
    function Get_GatewayEncryptedAuthCookieSize: LongWord; safecall;
    procedure Set_GatewayAuthLoginPage(const pbstrProxyAuthLoginPage: WideString); safecall;
    function Get_GatewayAuthLoginPage: WideString; safecall;
    property GatewayCredSourceCookie: LongWord read Get_GatewayCredSourceCookie write Set_GatewayCredSourceCookie;
    property GatewayAuthCookieServerAddr: WideString read Get_GatewayAuthCookieServerAddr write Set_GatewayAuthCookieServerAddr;
    property GatewayEncryptedAuthCookie: WideString read Get_GatewayEncryptedAuthCookie write Set_GatewayEncryptedAuthCookie;
    property GatewayEncryptedAuthCookieSize: LongWord read Get_GatewayEncryptedAuthCookieSize write Set_GatewayEncryptedAuthCookieSize;
    property GatewayAuthLoginPage: WideString read Get_GatewayAuthLoginPage write Set_GatewayAuthLoginPage;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientTransportSettings3Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {3D5B21AC-748D-41DE-8F30-E15169586BD4}
// *********************************************************************//
  IMsRdpClientTransportSettings3Disp = dispinterface
    ['{3D5B21AC-748D-41DE-8F30-E15169586BD4}']
    property GatewayCredSourceCookie: LongWord dispid 226;
    property GatewayAuthCookieServerAddr: WideString dispid 227;
    property GatewayEncryptedAuthCookie: WideString dispid 228;
    property GatewayEncryptedAuthCookieSize: LongWord dispid 229;
    property GatewayAuthLoginPage: WideString dispid 230;
    property GatewayCredSharing: LongWord dispid 222;
    property GatewayPreAuthRequirement: LongWord dispid 217;
    property GatewayPreAuthServerAddr: WideString dispid 218;
    property GatewaySupportUrl: WideString dispid 219;
    property GatewayEncryptedOtpCookie: WideString dispid 220;
    property GatewayEncryptedOtpCookieSize: LongWord dispid 221;
    property GatewayUsername: WideString dispid 223;
    property GatewayDomain: WideString dispid 224;
    property GatewayPassword: WideString writeonly dispid 225;
    property GatewayHostname: WideString dispid 210;
    property GatewayUsageMethod: LongWord dispid 211;
    property GatewayProfileUsageMethod: LongWord dispid 212;
    property GatewayCredsSource: LongWord dispid 213;
    property GatewayUserSelectedCredsSource: LongWord dispid 216;
    property GatewayIsSupported: Integer readonly dispid 214;
    property GatewayDefaultUsageMethod: LongWord readonly dispid 215;
  end;

// *********************************************************************//
// Interface: IMsRdpClientSecuredSettings2
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {25F2CE20-8B1D-4971-A7CD-549DAE201FC0}
// *********************************************************************//
  IMsRdpClientSecuredSettings2 = interface(IMsRdpClientSecuredSettings)
    ['{25F2CE20-8B1D-4971-A7CD-549DAE201FC0}']
    function Get_PCB: WideString; safecall;
    procedure Set_PCB(const bstrPCB: WideString); safecall;
    property PCB: WideString read Get_PCB write Set_PCB;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientSecuredSettings2Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {25F2CE20-8B1D-4971-A7CD-549DAE201FC0}
// *********************************************************************//
  IMsRdpClientSecuredSettings2Disp = dispinterface
    ['{25F2CE20-8B1D-4971-A7CD-549DAE201FC0}']
    property PCB: WideString dispid 6;
    property KeyboardHookMode: Integer dispid 4;
    property AudioRedirectionMode: Integer dispid 5;
    property StartProgram: WideString dispid 1;
    property WorkDir: WideString dispid 2;
    property FullScreen: Integer dispid 3;
  end;

// *********************************************************************//
// Interface: ITSRemoteProgram2
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {92C38A7D-241A-418C-9936-099872C9AF20}
// *********************************************************************//
  ITSRemoteProgram2 = interface(ITSRemoteProgram)
    ['{92C38A7D-241A-418C-9936-099872C9AF20}']
    procedure Set_RemoteApplicationName(const Param1: WideString); safecall;
    procedure Set_RemoteApplicationProgram(const Param1: WideString); safecall;
    procedure Set_RemoteApplicationArgs(const Param1: WideString); safecall;
    property RemoteApplicationName: WideString write Set_RemoteApplicationName;
    property RemoteApplicationProgram: WideString write Set_RemoteApplicationProgram;
    property RemoteApplicationArgs: WideString write Set_RemoteApplicationArgs;
  end;

// *********************************************************************//
// DispIntf:  ITSRemoteProgram2Disp
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {92C38A7D-241A-418C-9936-099872C9AF20}
// *********************************************************************//
  ITSRemoteProgram2Disp = dispinterface
    ['{92C38A7D-241A-418C-9936-099872C9AF20}']
    property RemoteApplicationName: WideString writeonly dispid 202;
    property RemoteApplicationProgram: WideString writeonly dispid 203;
    property RemoteApplicationArgs: WideString writeonly dispid 204;
    property RemoteProgramMode: WordBool dispid 200;
    procedure ServerStartProgram(const bstrExecutablePath: WideString; 
                                 const bstrFilePath: WideString; 
                                 const bstrWorkingDirectory: WideString; 
                                 vbExpandEnvVarInWorkingDirectoryOnServer: WordBool; 
                                 const bstrArguments: WideString; 
                                 vbExpandEnvVarInArgumentsOnServer: WordBool); dispid 201;
  end;

// *********************************************************************//
// Interface: IMsRdpPreferredRedirectionInfo
// Flags:     (0)
// GUID:      {FDD029F9-9574-4DEF-8529-64B521CCCAA4}
// *********************************************************************//
  IMsRdpPreferredRedirectionInfo = interface(IUnknown)
    ['{FDD029F9-9574-4DEF-8529-64B521CCCAA4}']
    function Set_UseRedirectionServerName(pVal: WordBool): HResult; stdcall;
    function Get_UseRedirectionServerName(out pVal: WordBool): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpExtendedSettings
// Flags:     (0)
// GUID:      {302D8188-0052-4807-806A-362B628F9AC5}
// *********************************************************************//
  IMsRdpExtendedSettings = interface(IUnknown)
    ['{302D8188-0052-4807-806A-362B628F9AC5}']
    function Set_Property_(const bstrPropertyName: WideString; const pValue: OleVariant): HResult; stdcall;
    function Get_Property_(const bstrPropertyName: WideString; out pValue: OleVariant): HResult; stdcall;
  end;

// *********************************************************************//
// Interface: IMsRdpClient8
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {4247E044-9271-43A9-BC49-E2AD9E855D62}
// *********************************************************************//
  IMsRdpClient8 = interface(IMsRdpClient7)
    ['{4247E044-9271-43A9-BC49-E2AD9E855D62}']
    procedure SendRemoteAction(actionType: RemoteSessionActionType); safecall;
    function Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8; safecall;
    function Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus; safecall;
    property AdvancedSettings9: IMsRdpClientAdvancedSettings8 read Get_AdvancedSettings9;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClient8Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {4247E044-9271-43A9-BC49-E2AD9E855D62}
// *********************************************************************//
  IMsRdpClient8Disp = dispinterface
    ['{4247E044-9271-43A9-BC49-E2AD9E855D62}']
    procedure SendRemoteAction(actionType: RemoteSessionActionType); dispid 700;
    property AdvancedSettings9: IMsRdpClientAdvancedSettings8 readonly dispid 701;
    function Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus; dispid 702;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 readonly dispid 600;
    property TransportSettings3: IMsRdpClientTransportSettings3 readonly dispid 601;
    function GetStatusText(statusCode: SYSUINT): WideString; dispid 602;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 readonly dispid 603;
    property RemoteProgram2: ITSRemoteProgram2 readonly dispid 604;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 readonly dispid 507;
    property TransportSettings2: IMsRdpClientTransportSettings2 readonly dispid 506;
    property TransportSettings: IMsRdpClientTransportSettings readonly dispid 500;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 readonly dispid 502;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString; dispid 503;
    property RemoteProgram: ITSRemoteProgram readonly dispid 504;
    property MsRdpClientShell: IMsRdpClientShell readonly dispid 505;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 readonly dispid 400;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 readonly dispid 300;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 readonly dispid 200;
    property ConnectedStatusText: WideString dispid 201;
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClientAdvancedSettings8
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {89ACB528-2557-4D16-8625-226A30E97E9A}
// *********************************************************************//
  IMsRdpClientAdvancedSettings8 = interface(IMsRdpClientAdvancedSettings7)
    ['{89ACB528-2557-4D16-8625-226A30E97E9A}']
    procedure Set_BandwidthDetection(pfAutodetect: WordBool); safecall;
    function Get_BandwidthDetection: WordBool; safecall;
    procedure Set_ClientProtocolSpec(pClientMode: ClientSpec); safecall;
    function Get_ClientProtocolSpec: ClientSpec; safecall;
    property BandwidthDetection: WordBool read Get_BandwidthDetection write Set_BandwidthDetection;
    property ClientProtocolSpec: ClientSpec read Get_ClientProtocolSpec write Set_ClientProtocolSpec;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientAdvancedSettings8Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {89ACB528-2557-4D16-8625-226A30E97E9A}
// *********************************************************************//
  IMsRdpClientAdvancedSettings8Disp = dispinterface
    ['{89ACB528-2557-4D16-8625-226A30E97E9A}']
    property BandwidthDetection: WordBool dispid 236;
    property ClientProtocolSpec: ClientSpec dispid 237;
    property AudioCaptureRedirectionMode: WordBool dispid 228;
    property VideoPlaybackMode: SYSUINT dispid 229;
    property EnableSuperPan: WordBool dispid 230;
    property SuperPanAccelerationFactor: LongWord dispid 231;
    property NegotiateSecurityLayer: WordBool dispid 232;
    property AudioQualityMode: SYSUINT dispid 233;
    property RedirectDirectX: WordBool dispid 234;
    property NetworkConnectionType: SYSUINT dispid 235;
    property RelativeMouseMode: WordBool dispid 221;
    property AuthenticationServiceClass: WideString dispid 222;
    property PCB: WideString dispid 223;
    property HotKeyFocusReleaseLeft: Integer dispid 224;
    property HotKeyFocusReleaseRight: Integer dispid 225;
    property EnableCredSspSupport: WordBool dispid 17;
    property AuthenticationType: SYSUINT readonly dispid 226;
    property ConnectToAdministerServer: WordBool dispid 227;
    property RedirectClipboard: WordBool dispid 213;
    property AudioRedirectionMode: SYSUINT dispid 215;
    property ConnectionBarShowPinButton: WordBool dispid 216;
    property PublicMode: WordBool dispid 217;
    property RedirectDevices: WordBool dispid 218;
    property RedirectPOSDevices: WordBool dispid 219;
    property BitmapVirtualCache32BppSize: Integer dispid 220;
    property AuthenticationLevel: SYSUINT dispid 212;
    property ConnectionBarShowMinimizeButton: WordBool dispid 210;
    property ConnectionBarShowRestoreButton: WordBool dispid 211;
    property CanAutoReconnect: WordBool readonly dispid 205;
    property EnableAutoReconnect: WordBool dispid 206;
    property MaxReconnectAttempts: Integer dispid 207;
    property SmoothScroll: Integer dispid 101;
    property AcceleratorPassthrough: Integer dispid 102;
    property ShadowBitmap: Integer dispid 103;
    property TransportType: Integer dispid 104;
    property SasSequence: Integer dispid 105;
    property EncryptionEnabled: Integer dispid 106;
    property DedicatedTerminal: Integer dispid 107;
    property RDPPort: Integer dispid 108;
    property EnableMouse: Integer dispid 109;
    property DisableCtrlAltDel: Integer dispid 110;
    property EnableWindowsKey: Integer dispid 111;
    property DoubleClickDetect: Integer dispid 112;
    property MaximizeShell: Integer dispid 113;
    property HotKeyFullScreen: Integer dispid 114;
    property HotKeyCtrlEsc: Integer dispid 115;
    property HotKeyAltEsc: Integer dispid 116;
    property HotKeyAltTab: Integer dispid 117;
    property HotKeyAltShiftTab: Integer dispid 118;
    property HotKeyAltSpace: Integer dispid 119;
    property HotKeyCtrlAltDel: Integer dispid 120;
    property orderDrawThreshold: Integer dispid 123;
    property BitmapCacheSize: Integer dispid 124;
    property BitmapVirtualCacheSize: Integer dispid 125;
    property ScaleBitmapCachesByBPP: Integer dispid 175;
    property NumBitmapCaches: Integer dispid 126;
    property CachePersistenceActive: Integer dispid 127;
    property PersistCacheDirectory: WideString writeonly dispid 138;
    property brushSupportLevel: Integer dispid 156;
    property minInputSendInterval: Integer dispid 157;
    property InputEventsAtOnce: Integer dispid 158;
    property maxEventCount: Integer dispid 159;
    property keepAliveInterval: Integer dispid 160;
    property shutdownTimeout: Integer dispid 163;
    property overallConnectionTimeout: Integer dispid 164;
    property singleConnectionTimeout: Integer dispid 165;
    property KeyboardType: Integer dispid 166;
    property KeyboardSubType: Integer dispid 167;
    property KeyboardFunctionKey: Integer dispid 168;
    property WinceFixedPalette: Integer dispid 169;
    property ConnectToServerConsole: WordBool dispid 178;
    property BitmapPersistence: Integer dispid 182;
    property MinutesToIdleTimeout: Integer dispid 183;
    property SmartSizing: WordBool dispid 184;
    property RdpdrLocalPrintingDocName: WideString dispid 185;
    property RdpdrClipCleanTempDirString: WideString dispid 201;
    property RdpdrClipPasteInfoString: WideString dispid 202;
    property ClearTextPassword: WideString writeonly dispid 186;
    property DisplayConnectionBar: WordBool dispid 187;
    property PinConnectionBar: WordBool dispid 188;
    property GrabFocusOnConnect: WordBool dispid 189;
    property LoadBalanceInfo: WideString dispid 190;
    property RedirectDrives: WordBool dispid 191;
    property RedirectPrinters: WordBool dispid 192;
    property RedirectPorts: WordBool dispid 193;
    property RedirectSmartCards: WordBool dispid 194;
    property BitmapVirtualCache16BppSize: Integer dispid 195;
    property BitmapVirtualCache24BppSize: Integer dispid 196;
    property PerformanceFlags: Integer dispid 200;
    property ConnectWithEndpoint: {NOT_OLEAUTO(POleVariant1)}OleVariant writeonly dispid 203;
    property NotifyTSPublicKey: WordBool dispid 204;
    property Compress: Integer dispid 121;
    property BitmapPeristence: Integer dispid 122;
    property allowBackgroundInput: Integer dispid 161;
    property KeyBoardLayoutStr: WideString writeonly dispid 162;
    property PluginDlls: WideString writeonly dispid 170;
    property IconFile: WideString writeonly dispid 171;
    property IconIndex: Integer writeonly dispid 172;
    property ContainerHandledFullScreen: Integer dispid 173;
    property DisableRdpdr: Integer dispid 174;
  end;

// *********************************************************************//
// Interface: IMsRdpClient9
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {28904001-04B6-436C-A55B-0AF1A0883DC9}
// *********************************************************************//
  IMsRdpClient9 = interface(IMsRdpClient8)
    ['{28904001-04B6-436C-A55B-0AF1A0883DC9}']
    function Get_TransportSettings4: IMsRdpClientTransportSettings4; safecall;
    procedure SyncSessionDisplaySettings; safecall;
    procedure UpdateSessionDisplaySettings(ulDesktopWidth: LongWord; ulDesktopHeight: LongWord; 
                                           ulPhysicalWidth: LongWord; ulPhysicalHeight: LongWord; 
                                           ulOrientation: LongWord; ulDesktopScaleFactor: LongWord; 
                                           ulDeviceScaleFactor: LongWord); safecall;
    procedure attachEvent(const eventName: WideString; const callback: IDispatch); safecall;
    procedure detachEvent(const eventName: WideString; const callback: IDispatch); safecall;
    property TransportSettings4: IMsRdpClientTransportSettings4 read Get_TransportSettings4;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClient9Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {28904001-04B6-436C-A55B-0AF1A0883DC9}
// *********************************************************************//
  IMsRdpClient9Disp = dispinterface
    ['{28904001-04B6-436C-A55B-0AF1A0883DC9}']
    property TransportSettings4: IMsRdpClientTransportSettings4 readonly dispid 800;
    procedure SyncSessionDisplaySettings; dispid 801;
    procedure UpdateSessionDisplaySettings(ulDesktopWidth: LongWord; ulDesktopHeight: LongWord; 
                                           ulPhysicalWidth: LongWord; ulPhysicalHeight: LongWord; 
                                           ulOrientation: LongWord; ulDesktopScaleFactor: LongWord; 
                                           ulDeviceScaleFactor: LongWord); dispid 802;
    procedure attachEvent(const eventName: WideString; const callback: IDispatch); dispid 803;
    procedure detachEvent(const eventName: WideString; const callback: IDispatch); dispid 804;
    procedure SendRemoteAction(actionType: RemoteSessionActionType); dispid 700;
    property AdvancedSettings9: IMsRdpClientAdvancedSettings8 readonly dispid 701;
    function Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus; dispid 702;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 readonly dispid 600;
    property TransportSettings3: IMsRdpClientTransportSettings3 readonly dispid 601;
    function GetStatusText(statusCode: SYSUINT): WideString; dispid 602;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 readonly dispid 603;
    property RemoteProgram2: ITSRemoteProgram2 readonly dispid 604;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 readonly dispid 507;
    property TransportSettings2: IMsRdpClientTransportSettings2 readonly dispid 506;
    property TransportSettings: IMsRdpClientTransportSettings readonly dispid 500;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 readonly dispid 502;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString; dispid 503;
    property RemoteProgram: ITSRemoteProgram readonly dispid 504;
    property MsRdpClientShell: IMsRdpClientShell readonly dispid 505;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 readonly dispid 400;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 readonly dispid 300;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 readonly dispid 200;
    property ConnectedStatusText: WideString dispid 201;
    property ColorDepth: Integer dispid 100;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings readonly dispid 101;
    property SecuredSettings2: IMsRdpClientSecuredSettings readonly dispid 102;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode readonly dispid 103;
    property FullScreen: WordBool dispid 104;
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer); dispid 35;
    function GetVirtualChannelOptions(const chanName: WideString): Integer; dispid 36;
    function RequestClose: ControlCloseStatus; dispid 37;
    property Server: WideString dispid 1;
    property Domain: WideString dispid 2;
    property UserName: WideString dispid 3;
    property DisconnectedText: WideString dispid 4;
    property ConnectingText: WideString dispid 5;
    property Connected: Smallint readonly dispid 6;
    property DesktopWidth: Integer dispid 12;
    property DesktopHeight: Integer dispid 13;
    property StartConnected: Integer dispid 16;
    property HorizontalScrollBarVisible: Integer readonly dispid 17;
    property VerticalScrollBarVisible: Integer readonly dispid 18;
    property FullScreenTitle: WideString writeonly dispid 19;
    property CipherStrength: Integer readonly dispid 20;
    property Version: WideString readonly dispid 21;
    property SecuredSettingsEnabled: Integer readonly dispid 22;
    property SecuredSettings: IMsTscSecuredSettings readonly dispid 97;
    property AdvancedSettings: IMsTscAdvancedSettings readonly dispid 98;
    property Debugger: IMsTscDebug readonly dispid 99;
    procedure Connect; dispid 30;
    procedure Disconnect; dispid 31;
    procedure CreateVirtualChannels(const newVal: WideString); dispid 33;
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString); dispid 34;
  end;

// *********************************************************************//
// Interface: IMsRdpClientTransportSettings4
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {011C3236-4D81-4515-9143-067AB630D299}
// *********************************************************************//
  IMsRdpClientTransportSettings4 = interface(IMsRdpClientTransportSettings3)
    ['{011C3236-4D81-4515-9143-067AB630D299}']
    procedure Set_GatewayBrokeringType(Param1: LongWord); safecall;
    property GatewayBrokeringType: LongWord write Set_GatewayBrokeringType;
  end;

// *********************************************************************//
// DispIntf:  IMsRdpClientTransportSettings4Disp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {011C3236-4D81-4515-9143-067AB630D299}
// *********************************************************************//
  IMsRdpClientTransportSettings4Disp = dispinterface
    ['{011C3236-4D81-4515-9143-067AB630D299}']
    property GatewayBrokeringType: LongWord writeonly dispid 231;
    property GatewayCredSourceCookie: LongWord dispid 226;
    property GatewayAuthCookieServerAddr: WideString dispid 227;
    property GatewayEncryptedAuthCookie: WideString dispid 228;
    property GatewayEncryptedAuthCookieSize: LongWord dispid 229;
    property GatewayAuthLoginPage: WideString dispid 230;
    property GatewayCredSharing: LongWord dispid 222;
    property GatewayPreAuthRequirement: LongWord dispid 217;
    property GatewayPreAuthServerAddr: WideString dispid 218;
    property GatewaySupportUrl: WideString dispid 219;
    property GatewayEncryptedOtpCookie: WideString dispid 220;
    property GatewayEncryptedOtpCookieSize: LongWord dispid 221;
    property GatewayUsername: WideString dispid 223;
    property GatewayDomain: WideString dispid 224;
    property GatewayPassword: WideString writeonly dispid 225;
    property GatewayHostname: WideString dispid 210;
    property GatewayUsageMethod: LongWord dispid 211;
    property GatewayProfileUsageMethod: LongWord dispid 212;
    property GatewayCredsSource: LongWord dispid 213;
    property GatewayUserSelectedCredsSource: LongWord dispid 216;
    property GatewayIsSupported: Integer readonly dispid 214;
    property GatewayDefaultUsageMethod: LongWord readonly dispid 215;
  end;

// *********************************************************************//
// DispIntf:  IRemoteDesktopClientEvents
// Flags:     (4096) Dispatchable
// GUID:      {079863B7-6D47-4105-8BFE-0CDCB360E67D}
// *********************************************************************//
  IRemoteDesktopClientEvents = dispinterface
    ['{079863B7-6D47-4105-8BFE-0CDCB360E67D}']
    procedure OnConnecting; dispid 750;
    procedure OnConnected; dispid 751;
    procedure OnLoginCompleted; dispid 752;
    procedure OnDisconnected(disconnectReason: Integer; ExtendedDisconnectReason: Integer; 
                             const disconnectErrorMessage: WideString); dispid 753;
    procedure OnStatusChanged(statusCode: Integer; const statusMessage: WideString); dispid 754;
    procedure OnAutoReconnecting(disconnectReason: Integer; ExtendedDisconnectReason: Integer; 
                                 const disconnectErrorMessage: WideString; 
                                 networkAvailable: WordBool; attemptCount: Integer; 
                                 maxAttemptCount: Integer); dispid 755;
    procedure OnAutoReconnected; dispid 756;
    procedure OnDialogDisplaying; dispid 757;
    procedure OnDialogDismissed; dispid 758;
    procedure OnNetworkStatusChanged(qualityLevel: LongWord; bandwidth: Integer; rtt: Integer); dispid 759;
    procedure OnAdminMessageReceived(const adminMessage: WideString); dispid 760;
    procedure OnKeyCombinationPressed(keyCombination: Integer); dispid 761;
    procedure OnRemoteDesktopSizeChanged(width: Integer; height: Integer); dispid 762;
    procedure OnTouchPointerCursorMoved(x: Integer; y: Integer); dispid 800;
  end;

// *********************************************************************//
// Interface: IRemoteDesktopClient
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {57D25668-625A-4905-BE4E-304CAA13F89C}
// *********************************************************************//
  IRemoteDesktopClient = interface(IDispatch)
    ['{57D25668-625A-4905-BE4E-304CAA13F89C}']
    procedure Connect; safecall;
    procedure Disconnect; safecall;
    procedure Reconnect(width: LongWord; height: LongWord); safecall;
    function Get_Settings: IRemoteDesktopClientSettings; safecall;
    function Get_Actions: IRemoteDesktopClientActions; safecall;
    function Get_TouchPointer: IRemoteDesktopClientTouchPointer; safecall;
    procedure DeleteSavedCredentials(const serverName: WideString); safecall;
    procedure UpdateSessionDisplaySettings(width: LongWord; height: LongWord); safecall;
    procedure attachEvent(const eventName: WideString; const callback: IDispatch); safecall;
    procedure detachEvent(const eventName: WideString; const callback: IDispatch); safecall;
    property Settings: IRemoteDesktopClientSettings read Get_Settings;
    property Actions: IRemoteDesktopClientActions read Get_Actions;
    property TouchPointer: IRemoteDesktopClientTouchPointer read Get_TouchPointer;
  end;

// *********************************************************************//
// DispIntf:  IRemoteDesktopClientDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {57D25668-625A-4905-BE4E-304CAA13F89C}
// *********************************************************************//
  IRemoteDesktopClientDisp = dispinterface
    ['{57D25668-625A-4905-BE4E-304CAA13F89C}']
    procedure Connect; dispid 701;
    procedure Disconnect; dispid 702;
    procedure Reconnect(width: LongWord; height: LongWord); dispid 703;
    property Settings: IRemoteDesktopClientSettings readonly dispid 710;
    property Actions: IRemoteDesktopClientActions readonly dispid 711;
    property TouchPointer: IRemoteDesktopClientTouchPointer readonly dispid 712;
    procedure DeleteSavedCredentials(const serverName: WideString); dispid 704;
    procedure UpdateSessionDisplaySettings(width: LongWord; height: LongWord); dispid 705;
    procedure attachEvent(const eventName: WideString; const callback: IDispatch); dispid 706;
    procedure detachEvent(const eventName: WideString; const callback: IDispatch); dispid 707;
  end;

// *********************************************************************//
// Interface: IRemoteDesktopClientSettings
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {48A0F2A7-2713-431F-BBAC-6F4558E7D64D}
// *********************************************************************//
  IRemoteDesktopClientSettings = interface(IDispatch)
    ['{48A0F2A7-2713-431F-BBAC-6F4558E7D64D}']
    procedure ApplySettings(const RdpFileContents: WideString); safecall;
    function RetrieveSettings: WideString; safecall;
    function GetRdpProperty(const propertyName: WideString): OleVariant; safecall;
    procedure SetRdpProperty(const propertyName: WideString; Value: OleVariant); safecall;
  end;

// *********************************************************************//
// DispIntf:  IRemoteDesktopClientSettingsDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {48A0F2A7-2713-431F-BBAC-6F4558E7D64D}
// *********************************************************************//
  IRemoteDesktopClientSettingsDisp = dispinterface
    ['{48A0F2A7-2713-431F-BBAC-6F4558E7D64D}']
    procedure ApplySettings(const RdpFileContents: WideString); dispid 722;
    function RetrieveSettings: WideString; dispid 723;
    function GetRdpProperty(const propertyName: WideString): OleVariant; dispid 721;
    procedure SetRdpProperty(const propertyName: WideString; Value: OleVariant); dispid 720;
  end;

// *********************************************************************//
// Interface: IRemoteDesktopClientActions
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {7D54BC4E-1028-45D4-8B0A-B9B6BFFBA176}
// *********************************************************************//
  IRemoteDesktopClientActions = interface(IDispatch)
    ['{7D54BC4E-1028-45D4-8B0A-B9B6BFFBA176}']
    procedure SuspendScreenUpdates; safecall;
    procedure ResumeScreenUpdates; safecall;
    procedure ExecuteRemoteAction(remoteAction: RemoteActionType); safecall;
    function GetSnapshot(snapshotEncoding: SnapshotEncodingType; 
                         snapshotFormat: SnapshotFormatType; snapshotWidth: LongWord; 
                         snapshotHeight: LongWord): WideString; safecall;
  end;

// *********************************************************************//
// DispIntf:  IRemoteDesktopClientActionsDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {7D54BC4E-1028-45D4-8B0A-B9B6BFFBA176}
// *********************************************************************//
  IRemoteDesktopClientActionsDisp = dispinterface
    ['{7D54BC4E-1028-45D4-8B0A-B9B6BFFBA176}']
    procedure SuspendScreenUpdates; dispid 730;
    procedure ResumeScreenUpdates; dispid 731;
    procedure ExecuteRemoteAction(remoteAction: RemoteActionType); dispid 732;
    function GetSnapshot(snapshotEncoding: SnapshotEncodingType; 
                         snapshotFormat: SnapshotFormatType; snapshotWidth: LongWord; 
                         snapshotHeight: LongWord): WideString; dispid 733;
  end;

// *********************************************************************//
// Interface: IRemoteDesktopClientTouchPointer
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {260EC22D-8CBC-44B5-9E88-2A37F6C93AE9}
// *********************************************************************//
  IRemoteDesktopClientTouchPointer = interface(IDispatch)
    ['{260EC22D-8CBC-44B5-9E88-2A37F6C93AE9}']
    procedure Set_Enabled(Enabled: WordBool); safecall;
    function Get_Enabled: WordBool; safecall;
    procedure Set_EventsEnabled(EventsEnabled: WordBool); safecall;
    function Get_EventsEnabled: WordBool; safecall;
    procedure Set_PointerSpeed(PointerSpeed: LongWord); safecall;
    function Get_PointerSpeed: LongWord; safecall;
    property Enabled: WordBool read Get_Enabled write Set_Enabled;
    property EventsEnabled: WordBool read Get_EventsEnabled write Set_EventsEnabled;
    property PointerSpeed: LongWord read Get_PointerSpeed write Set_PointerSpeed;
  end;

// *********************************************************************//
// DispIntf:  IRemoteDesktopClientTouchPointerDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {260EC22D-8CBC-44B5-9E88-2A37F6C93AE9}
// *********************************************************************//
  IRemoteDesktopClientTouchPointerDisp = dispinterface
    ['{260EC22D-8CBC-44B5-9E88-2A37F6C93AE9}']
    property Enabled: WordBool dispid 740;
    property EventsEnabled: WordBool dispid 741;
    property PointerSpeed: LongWord dispid 742;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsTscAxNotSafeForScripting
// Help String      : Microsoft Terminal Services Client Control - version 1
// Default Interface: IMsTscAx
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsTscAxNotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsTscAxNotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                 const data: WideString) of object;
  TMsTscAxNotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsTscAxNotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsTscAxNotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                     height: Integer) of object;
  TMsTscAxNotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsTscAxNotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                              attemptCount: Integer) of object;
  TMsTscAxNotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                 lError: RemoteProgramResult; 
                                                                                 vbIsExecutable: WordBool) of object;
  TMsTscAxNotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                    uDisplayInformation: LongWord) of object;
  TMsTscAxNotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                   var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                   windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsTscAxNotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsTscAxNotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsTscAxNotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsTscAxNotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsTscAxNotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsTscAxNotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                  bandwidth: Integer; 
                                                                                  rtt: Integer) of object;
  TMsTscAxNotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                               networkAvailable: WordBool; 
                                                                               attemptCount: Integer; 
                                                                               maxAttemptCount: Integer) of object;

  TMsTscAxNotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsTscAxNotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsTscAxNotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsTscAxNotSafeForScriptingOnFatalError;
    FOnWarning: TMsTscAxNotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsTscAxNotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsTscAxNotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsTscAxNotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsTscAxNotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsTscAxNotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsTscAxNotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsTscAxNotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsTscAxNotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsTscAxNotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsTscAxNotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsTscAxNotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsTscAxNotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsTscAxNotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsTscAx;
    function  GetControlInterface: IMsTscAx;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    property  ControlInterface: IMsTscAx read GetControlInterface;
    property  DefaultInterface: IMsTscAx read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsTscAxNotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsTscAxNotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsTscAxNotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsTscAxNotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsTscAxNotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsTscAxNotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsTscAxNotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsTscAxNotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsTscAxNotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsTscAxNotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsTscAxNotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsTscAxNotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsTscAxNotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsTscAxNotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsTscAxNotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsTscAxNotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsTscAxNotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;

// *********************************************************************//
// The Class CoMsTscAx provides a Create and CreateRemote method to          
// create instances of the default interface IMsTscAx exposed by              
// the CoClass MsTscAx. The functions are intended to be used by             
// clients wishing to automate the CoClass objects exposed by the         
// server of this typelibrary.                                            
// *********************************************************************//
  CoMsTscAx = class
    class function Create: IMsTscAx;
    class function CreateRemote(const MachineName: string): IMsTscAx;
  end;

  TMsTscAxOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsTscAxOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                              const data: WideString) of object;
  TMsTscAxOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsTscAxOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsTscAxOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; height: Integer) of object;
  TMsTscAxOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsTscAxOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                           attemptCount: Integer) of object;
  TMsTscAxOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                              lError: RemoteProgramResult; 
                                                              vbIsExecutable: WordBool) of object;
  TMsTscAxOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                 uDisplayInformation: LongWord) of object;
  TMsTscAxOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsTscAxOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsTscAxOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsTscAxOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsTscAxOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsTscAxOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsTscAxOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                               bandwidth: Integer; rtt: Integer) of object;
  TMsTscAxOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                            networkAvailable: WordBool; 
                                                            attemptCount: Integer; 
                                                            maxAttemptCount: Integer) of object;


// *********************************************************************//
// OLE Server Proxy class declaration
// Server Object    : TMsTscAx
// Help String      : Microsoft Terminal Services Client Control (redistributable) - version 1
// Default Interface: IMsTscAx
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (0)
// *********************************************************************//
  TMsTscAx = class(TOleServer)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsTscAxOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsTscAxOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsTscAxOnFatalError;
    FOnWarning: TMsTscAxOnWarning;
    FOnRemoteDesktopSizeChange: TMsTscAxOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsTscAxOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsTscAxOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsTscAxOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsTscAxOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsTscAxOnRemoteWindowDisplayed;
    FOnLogonError: TMsTscAxOnLogonError;
    FOnFocusReleased: TMsTscAxOnFocusReleased;
    FOnUserNameAcquired: TMsTscAxOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsTscAxOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsTscAxOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsTscAxOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsTscAxOnAutoReconnecting2;
    FIntf: IMsTscAx;
    function GetDefaultInterface: IMsTscAx;
  protected
    procedure InitServerData; override;
    procedure InvokeEvent(DispID: TDispID; var Params: TVariantArray); override;
    procedure Set_Server(const pServer: WideString);
    function Get_Server: WideString;
    procedure Set_Domain(const pDomain: WideString);
    function Get_Domain: WideString;
    procedure Set_UserName(const pUserName: WideString);
    function Get_UserName: WideString;
    procedure Set_DisconnectedText(const pDisconnectedText: WideString);
    function Get_DisconnectedText: WideString;
    procedure Set_ConnectingText(const pConnectingText: WideString);
    function Get_ConnectingText: WideString;
    function Get_Connected: Smallint;
    procedure Set_DesktopWidth(pVal: Integer);
    function Get_DesktopWidth: Integer;
    procedure Set_DesktopHeight(pVal: Integer);
    function Get_DesktopHeight: Integer;
    procedure Set_StartConnected(pfStartConnected: Integer);
    function Get_StartConnected: Integer;
    function Get_HorizontalScrollBarVisible: Integer;
    function Get_VerticalScrollBarVisible: Integer;
    procedure Set_FullScreenTitle(const Param1: WideString);
    function Get_CipherStrength: Integer;
    function Get_Version: WideString;
    function Get_SecuredSettingsEnabled: Integer;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
  public
    constructor Create(AOwner: TComponent); override;
    destructor  Destroy; override;
    procedure Connect; override;
    procedure ConnectTo(svrIntf: IMsTscAx);
    procedure Disconnect; override;
    procedure Connect1;
    procedure Disconnect1;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    property DefaultInterface: IMsTscAx read GetDefaultInterface;
    property Connected: Smallint read Get_Connected;
    property HorizontalScrollBarVisible: Integer read Get_HorizontalScrollBarVisible;
    property VerticalScrollBarVisible: Integer read Get_VerticalScrollBarVisible;
    property FullScreenTitle: WideString write Set_FullScreenTitle;
    property CipherStrength: Integer read Get_CipherStrength;
    property Version: WideString read Get_Version;
    property SecuredSettingsEnabled: Integer read Get_SecuredSettingsEnabled;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property Server: WideString read Get_Server write Set_Server;
    property Domain: WideString read Get_Domain write Set_Domain;
    property UserName: WideString read Get_UserName write Set_UserName;
    property DisconnectedText: WideString read Get_DisconnectedText write Set_DisconnectedText;
    property ConnectingText: WideString read Get_ConnectingText write Set_ConnectingText;
    property DesktopWidth: Integer read Get_DesktopWidth write Set_DesktopWidth;
    property DesktopHeight: Integer read Get_DesktopHeight write Set_DesktopHeight;
    property StartConnected: Integer read Get_StartConnected write Set_StartConnected;
  published
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsTscAxOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsTscAxOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsTscAxOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsTscAxOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsTscAxOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsTscAxOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsTscAxOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsTscAxOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsTscAxOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsTscAxOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsTscAxOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsTscAxOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsTscAxOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsTscAxOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsTscAxOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsTscAxOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsTscAxOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClientNotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 2
// Default Interface: IMsRdpClient
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClientNotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClientNotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                     const data: WideString) of object;
  TMsRdpClientNotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClientNotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClientNotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                         height: Integer) of object;
  TMsRdpClientNotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClientNotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                  attemptCount: Integer) of object;
  TMsRdpClientNotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                     lError: RemoteProgramResult; 
                                                                                     vbIsExecutable: WordBool) of object;
  TMsRdpClientNotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        uDisplayInformation: LongWord) of object;
  TMsRdpClientNotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                       var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                       windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClientNotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClientNotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClientNotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClientNotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClientNotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClientNotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                      bandwidth: Integer; 
                                                                                      rtt: Integer) of object;
  TMsRdpClientNotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   networkAvailable: WordBool; 
                                                                                   attemptCount: Integer; 
                                                                                   maxAttemptCount: Integer) of object;

  TMsRdpClientNotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClientNotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClientNotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClientNotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClientNotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClientNotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClientNotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClientNotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClientNotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClientNotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClientNotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClientNotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClientNotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClientNotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClientNotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClientNotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClientNotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClientNotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient;
    function  GetControlInterface: IMsRdpClient;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient read GetControlInterface;
    property  DefaultInterface: IMsRdpClient read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClientNotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClientNotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClientNotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClientNotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClientNotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClientNotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClientNotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClientNotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClientNotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClientNotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClientNotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClientNotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClientNotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClientNotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClientNotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClientNotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClientNotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;

// *********************************************************************//
// The Class CoMsRdpClient provides a Create and CreateRemote method to          
// create instances of the default interface IMsRdpClient exposed by              
// the CoClass MsRdpClient. The functions are intended to be used by             
// clients wishing to automate the CoClass objects exposed by the         
// server of this typelibrary.                                            
// *********************************************************************//
  CoMsRdpClient = class
    class function Create: IMsRdpClient;
    class function CreateRemote(const MachineName: string): IMsRdpClient;
  end;

  TMsRdpClientOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClientOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                  const data: WideString) of object;
  TMsRdpClientOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClientOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClientOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                      height: Integer) of object;
  TMsRdpClientOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClientOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                               attemptCount: Integer) of object;
  TMsRdpClientOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                  lError: RemoteProgramResult; 
                                                                  vbIsExecutable: WordBool) of object;
  TMsRdpClientOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     uDisplayInformation: LongWord) of object;
  TMsRdpClientOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                    var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                    windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClientOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClientOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClientOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClientOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClientOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClientOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                   bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClientOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                networkAvailable: WordBool; 
                                                                attemptCount: Integer; 
                                                                maxAttemptCount: Integer) of object;


// *********************************************************************//
// OLE Server Proxy class declaration
// Server Object    : TMsRdpClient
// Help String      : Microsoft RDP Client Control (redistributable) - version 2
// Default Interface: IMsRdpClient
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (0)
// *********************************************************************//
  TMsRdpClient = class(TOleServer)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClientOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClientOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClientOnFatalError;
    FOnWarning: TMsRdpClientOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClientOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClientOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClientOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClientOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClientOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClientOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClientOnLogonError;
    FOnFocusReleased: TMsRdpClientOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClientOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClientOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClientOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClientOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClientOnAutoReconnecting2;
    FIntf: IMsRdpClient;
    function GetDefaultInterface: IMsRdpClient;
  protected
    procedure InitServerData; override;
    procedure InvokeEvent(DispID: TDispID; var Params: TVariantArray); override;
    procedure Set_Server(const pServer: WideString);
    function Get_Server: WideString;
    procedure Set_Domain(const pDomain: WideString);
    function Get_Domain: WideString;
    procedure Set_UserName(const pUserName: WideString);
    function Get_UserName: WideString;
    procedure Set_DisconnectedText(const pDisconnectedText: WideString);
    function Get_DisconnectedText: WideString;
    procedure Set_ConnectingText(const pConnectingText: WideString);
    function Get_ConnectingText: WideString;
    function Get_Connected: Smallint;
    procedure Set_DesktopWidth(pVal: Integer);
    function Get_DesktopWidth: Integer;
    procedure Set_DesktopHeight(pVal: Integer);
    function Get_DesktopHeight: Integer;
    procedure Set_StartConnected(pfStartConnected: Integer);
    function Get_StartConnected: Integer;
    function Get_HorizontalScrollBarVisible: Integer;
    function Get_VerticalScrollBarVisible: Integer;
    procedure Set_FullScreenTitle(const Param1: WideString);
    function Get_CipherStrength: Integer;
    function Get_Version: WideString;
    function Get_SecuredSettingsEnabled: Integer;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    procedure Set_ColorDepth(pcolorDepth: Integer);
    function Get_ColorDepth: Integer;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_ExtendedDisconnectReason: ExtendedDisconnectReasonCode;
    procedure Set_FullScreen(pfFullScreen: WordBool);
    function Get_FullScreen: WordBool;
  public
    constructor Create(AOwner: TComponent); override;
    destructor  Destroy; override;
    procedure Connect; override;
    procedure ConnectTo(svrIntf: IMsRdpClient);
    procedure Disconnect; override;
    procedure Connect1;
    procedure Disconnect1;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property DefaultInterface: IMsRdpClient read GetDefaultInterface;
    property Connected: Smallint read Get_Connected;
    property HorizontalScrollBarVisible: Integer read Get_HorizontalScrollBarVisible;
    property VerticalScrollBarVisible: Integer read Get_VerticalScrollBarVisible;
    property FullScreenTitle: WideString write Set_FullScreenTitle;
    property CipherStrength: Integer read Get_CipherStrength;
    property Version: WideString read Get_Version;
    property SecuredSettingsEnabled: Integer read Get_SecuredSettingsEnabled;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: ExtendedDisconnectReasonCode read Get_ExtendedDisconnectReason;
    property Server: WideString read Get_Server write Set_Server;
    property Domain: WideString read Get_Domain write Set_Domain;
    property UserName: WideString read Get_UserName write Set_UserName;
    property DisconnectedText: WideString read Get_DisconnectedText write Set_DisconnectedText;
    property ConnectingText: WideString read Get_ConnectingText write Set_ConnectingText;
    property DesktopWidth: Integer read Get_DesktopWidth write Set_DesktopWidth;
    property DesktopHeight: Integer read Get_DesktopHeight write Set_DesktopHeight;
    property StartConnected: Integer read Get_StartConnected write Set_StartConnected;
    property ColorDepth: Integer read Get_ColorDepth write Set_ColorDepth;
    property FullScreen: WordBool read Get_FullScreen write Set_FullScreen;
  published
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClientOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClientOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClientOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClientOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClientOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClientOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClientOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClientOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClientOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClientOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClientOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClientOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClientOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClientOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClientOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClientOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClientOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient2NotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 3
// Default Interface: IMsRdpClient2
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient2NotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient2NotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                      const data: WideString) of object;
  TMsRdpClient2NotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient2NotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient2NotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                          height: Integer) of object;
  TMsRdpClient2NotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient2NotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   attemptCount: Integer) of object;
  TMsRdpClient2NotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                      lError: RemoteProgramResult; 
                                                                                      vbIsExecutable: WordBool) of object;
  TMsRdpClient2NotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                         uDisplayInformation: LongWord) of object;
  TMsRdpClient2NotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                        windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient2NotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient2NotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient2NotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient2NotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient2NotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient2NotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                       bandwidth: Integer; 
                                                                                       rtt: Integer) of object;
  TMsRdpClient2NotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                    networkAvailable: WordBool; 
                                                                                    attemptCount: Integer; 
                                                                                    maxAttemptCount: Integer) of object;

  TMsRdpClient2NotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient2NotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient2NotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient2NotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClient2NotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient2NotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient2NotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient2NotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient2NotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient2NotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient2NotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient2NotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClient2NotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient2NotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient2NotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient2NotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient2NotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient2NotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient2;
    function  GetControlInterface: IMsRdpClient2;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient2 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient2 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient2NotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient2NotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient2NotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient2NotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient2NotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient2NotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient2NotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient2NotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient2NotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient2NotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient2NotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient2NotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient2NotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient2NotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient2NotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient2NotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient2NotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient2
// Help String      : Microsoft RDP Client Control (redistributable) - version 3
// Default Interface: IMsRdpClient2
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient2OnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient2OnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                   const data: WideString) of object;
  TMsRdpClient2OnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient2OnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient2OnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                       height: Integer) of object;
  TMsRdpClient2OnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient2OnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                attemptCount: Integer) of object;
  TMsRdpClient2OnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                   lError: RemoteProgramResult; 
                                                                   vbIsExecutable: WordBool) of object;
  TMsRdpClient2OnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      uDisplayInformation: LongWord) of object;
  TMsRdpClient2OnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                     windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient2OnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient2OnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient2OnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient2OnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient2OnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient2OnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                    bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClient2OnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 networkAvailable: WordBool; 
                                                                 attemptCount: Integer; 
                                                                 maxAttemptCount: Integer) of object;

  TMsRdpClient2 = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient2OnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient2OnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient2OnFatalError;
    FOnWarning: TMsRdpClient2OnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient2OnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient2OnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient2OnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient2OnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient2OnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient2OnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient2OnLogonError;
    FOnFocusReleased: TMsRdpClient2OnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient2OnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient2OnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient2OnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient2OnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient2OnAutoReconnecting2;
    FIntf: IMsRdpClient2;
    function  GetControlInterface: IMsRdpClient2;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient2 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient2 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient2OnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient2OnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient2OnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient2OnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient2OnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient2OnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient2OnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient2OnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient2OnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient2OnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient2OnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient2OnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient2OnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient2OnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient2OnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient2OnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient2OnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient2a
// Help String      : Microsoft RDP Client Control (redistributable) - version 3a
// Default Interface: IMsRdpClient2
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient2aOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient2aOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                    const data: WideString) of object;
  TMsRdpClient2aOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient2aOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient2aOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                        height: Integer) of object;
  TMsRdpClient2aOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient2aOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 attemptCount: Integer) of object;
  TMsRdpClient2aOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                    lError: RemoteProgramResult; 
                                                                    vbIsExecutable: WordBool) of object;
  TMsRdpClient2aOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                       uDisplayInformation: LongWord) of object;
  TMsRdpClient2aOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                      windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient2aOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient2aOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient2aOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient2aOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient2aOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient2aOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                     bandwidth: Integer; 
                                                                     rtt: Integer) of object;
  TMsRdpClient2aOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                  networkAvailable: WordBool; 
                                                                  attemptCount: Integer; 
                                                                  maxAttemptCount: Integer) of object;

  TMsRdpClient2a = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient2aOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient2aOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient2aOnFatalError;
    FOnWarning: TMsRdpClient2aOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient2aOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient2aOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient2aOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient2aOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient2aOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient2aOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient2aOnLogonError;
    FOnFocusReleased: TMsRdpClient2aOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient2aOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient2aOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient2aOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient2aOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient2aOnAutoReconnecting2;
    FIntf: IMsRdpClient2;
    function  GetControlInterface: IMsRdpClient2;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient2 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient2 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient2aOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient2aOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient2aOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient2aOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient2aOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient2aOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient2aOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient2aOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient2aOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient2aOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient2aOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient2aOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient2aOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient2aOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient2aOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient2aOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient2aOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient3NotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 4
// Default Interface: IMsRdpClient3
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient3NotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient3NotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                      const data: WideString) of object;
  TMsRdpClient3NotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient3NotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient3NotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                          height: Integer) of object;
  TMsRdpClient3NotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient3NotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   attemptCount: Integer) of object;
  TMsRdpClient3NotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                      lError: RemoteProgramResult; 
                                                                                      vbIsExecutable: WordBool) of object;
  TMsRdpClient3NotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                         uDisplayInformation: LongWord) of object;
  TMsRdpClient3NotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                        windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient3NotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient3NotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient3NotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient3NotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient3NotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient3NotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                       bandwidth: Integer; 
                                                                                       rtt: Integer) of object;
  TMsRdpClient3NotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                    networkAvailable: WordBool; 
                                                                                    attemptCount: Integer; 
                                                                                    maxAttemptCount: Integer) of object;

  TMsRdpClient3NotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient3NotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient3NotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient3NotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClient3NotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient3NotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient3NotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient3NotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient3NotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient3NotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient3NotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient3NotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClient3NotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient3NotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient3NotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient3NotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient3NotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient3NotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient3;
    function  GetControlInterface: IMsRdpClient3;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient3 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient3 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient3NotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient3NotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient3NotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient3NotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient3NotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient3NotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient3NotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient3NotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient3NotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient3NotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient3NotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient3NotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient3NotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient3NotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient3NotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient3NotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient3NotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient3
// Help String      : Microsoft RDP Client Control (redistributable) - version 4
// Default Interface: IMsRdpClient3
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient3OnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient3OnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                   const data: WideString) of object;
  TMsRdpClient3OnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient3OnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient3OnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                       height: Integer) of object;
  TMsRdpClient3OnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient3OnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                attemptCount: Integer) of object;
  TMsRdpClient3OnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                   lError: RemoteProgramResult; 
                                                                   vbIsExecutable: WordBool) of object;
  TMsRdpClient3OnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      uDisplayInformation: LongWord) of object;
  TMsRdpClient3OnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                     windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient3OnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient3OnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient3OnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient3OnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient3OnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient3OnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                    bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClient3OnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 networkAvailable: WordBool; 
                                                                 attemptCount: Integer; 
                                                                 maxAttemptCount: Integer) of object;

  TMsRdpClient3 = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient3OnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient3OnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient3OnFatalError;
    FOnWarning: TMsRdpClient3OnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient3OnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient3OnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient3OnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient3OnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient3OnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient3OnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient3OnLogonError;
    FOnFocusReleased: TMsRdpClient3OnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient3OnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient3OnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient3OnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient3OnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient3OnAutoReconnecting2;
    FIntf: IMsRdpClient3;
    function  GetControlInterface: IMsRdpClient3;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient3 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient3 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient3OnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient3OnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient3OnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient3OnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient3OnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient3OnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient3OnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient3OnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient3OnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient3OnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient3OnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient3OnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient3OnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient3OnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient3OnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient3OnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient3OnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient3a
// Help String      : Microsoft RDP Client Control (redistributable) - version 4a
// Default Interface: IMsRdpClient3
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient3aOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient3aOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                    const data: WideString) of object;
  TMsRdpClient3aOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient3aOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient3aOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                        height: Integer) of object;
  TMsRdpClient3aOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient3aOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 attemptCount: Integer) of object;
  TMsRdpClient3aOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                    lError: RemoteProgramResult; 
                                                                    vbIsExecutable: WordBool) of object;
  TMsRdpClient3aOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                       uDisplayInformation: LongWord) of object;
  TMsRdpClient3aOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                      windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient3aOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient3aOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient3aOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient3aOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient3aOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient3aOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                     bandwidth: Integer; 
                                                                     rtt: Integer) of object;
  TMsRdpClient3aOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                  networkAvailable: WordBool; 
                                                                  attemptCount: Integer; 
                                                                  maxAttemptCount: Integer) of object;

  TMsRdpClient3a = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient3aOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient3aOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient3aOnFatalError;
    FOnWarning: TMsRdpClient3aOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient3aOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient3aOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient3aOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient3aOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient3aOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient3aOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient3aOnLogonError;
    FOnFocusReleased: TMsRdpClient3aOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient3aOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient3aOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient3aOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient3aOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient3aOnAutoReconnecting2;
    FIntf: IMsRdpClient3;
    function  GetControlInterface: IMsRdpClient3;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient3 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient3 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient3aOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient3aOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient3aOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient3aOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient3aOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient3aOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient3aOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient3aOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient3aOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient3aOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient3aOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient3aOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient3aOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient3aOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient3aOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient3aOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient3aOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient4NotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 5
// Default Interface: IMsRdpClient4
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient4NotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient4NotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                      const data: WideString) of object;
  TMsRdpClient4NotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient4NotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient4NotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                          height: Integer) of object;
  TMsRdpClient4NotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient4NotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   attemptCount: Integer) of object;
  TMsRdpClient4NotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                      lError: RemoteProgramResult; 
                                                                                      vbIsExecutable: WordBool) of object;
  TMsRdpClient4NotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                         uDisplayInformation: LongWord) of object;
  TMsRdpClient4NotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                        windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient4NotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient4NotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient4NotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient4NotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient4NotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient4NotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                       bandwidth: Integer; 
                                                                                       rtt: Integer) of object;
  TMsRdpClient4NotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                    networkAvailable: WordBool; 
                                                                                    attemptCount: Integer; 
                                                                                    maxAttemptCount: Integer) of object;

  TMsRdpClient4NotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient4NotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient4NotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient4NotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClient4NotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient4NotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient4NotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient4NotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient4NotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient4NotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient4NotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient4NotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClient4NotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient4NotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient4NotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient4NotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient4NotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient4NotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient4;
    function  GetControlInterface: IMsRdpClient4;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient4 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient4 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient4NotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient4NotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient4NotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient4NotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient4NotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient4NotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient4NotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient4NotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient4NotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient4NotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient4NotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient4NotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient4NotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient4NotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient4NotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient4NotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient4NotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient4
// Help String      : Microsoft RDP Client Control (redistributable) - version 5
// Default Interface: IMsRdpClient4
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient4OnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient4OnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                   const data: WideString) of object;
  TMsRdpClient4OnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient4OnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient4OnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                       height: Integer) of object;
  TMsRdpClient4OnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient4OnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                attemptCount: Integer) of object;
  TMsRdpClient4OnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                   lError: RemoteProgramResult; 
                                                                   vbIsExecutable: WordBool) of object;
  TMsRdpClient4OnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      uDisplayInformation: LongWord) of object;
  TMsRdpClient4OnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                     windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient4OnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient4OnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient4OnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient4OnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient4OnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient4OnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                    bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClient4OnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 networkAvailable: WordBool; 
                                                                 attemptCount: Integer; 
                                                                 maxAttemptCount: Integer) of object;

  TMsRdpClient4 = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient4OnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient4OnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient4OnFatalError;
    FOnWarning: TMsRdpClient4OnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient4OnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient4OnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient4OnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient4OnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient4OnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient4OnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient4OnLogonError;
    FOnFocusReleased: TMsRdpClient4OnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient4OnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient4OnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient4OnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient4OnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient4OnAutoReconnecting2;
    FIntf: IMsRdpClient4;
    function  GetControlInterface: IMsRdpClient4;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient4 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient4 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient4OnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient4OnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient4OnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient4OnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient4OnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient4OnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient4OnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient4OnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient4OnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient4OnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient4OnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient4OnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient4OnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient4OnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient4OnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient4OnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient4OnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient4a
// Help String      : Microsoft RDP Client Control (redistributable) - version 5a
// Default Interface: IMsRdpClient4
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient4aOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient4aOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                    const data: WideString) of object;
  TMsRdpClient4aOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient4aOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient4aOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                        height: Integer) of object;
  TMsRdpClient4aOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient4aOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 attemptCount: Integer) of object;
  TMsRdpClient4aOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                    lError: RemoteProgramResult; 
                                                                    vbIsExecutable: WordBool) of object;
  TMsRdpClient4aOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                       uDisplayInformation: LongWord) of object;
  TMsRdpClient4aOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                      windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient4aOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient4aOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient4aOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient4aOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient4aOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient4aOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                     bandwidth: Integer; 
                                                                     rtt: Integer) of object;
  TMsRdpClient4aOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                  networkAvailable: WordBool; 
                                                                  attemptCount: Integer; 
                                                                  maxAttemptCount: Integer) of object;

  TMsRdpClient4a = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient4aOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient4aOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient4aOnFatalError;
    FOnWarning: TMsRdpClient4aOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient4aOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient4aOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient4aOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient4aOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient4aOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient4aOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient4aOnLogonError;
    FOnFocusReleased: TMsRdpClient4aOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient4aOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient4aOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient4aOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient4aOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient4aOnAutoReconnecting2;
    FIntf: IMsRdpClient4;
    function  GetControlInterface: IMsRdpClient4;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    property  ControlInterface: IMsRdpClient4 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient4 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient4aOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient4aOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient4aOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient4aOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient4aOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient4aOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient4aOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient4aOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient4aOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient4aOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient4aOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient4aOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient4aOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient4aOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient4aOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient4aOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient4aOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient5NotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 6
// Default Interface: IMsRdpClient5
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient5NotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient5NotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                      const data: WideString) of object;
  TMsRdpClient5NotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient5NotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient5NotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                          height: Integer) of object;
  TMsRdpClient5NotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient5NotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   attemptCount: Integer) of object;
  TMsRdpClient5NotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                      lError: RemoteProgramResult; 
                                                                                      vbIsExecutable: WordBool) of object;
  TMsRdpClient5NotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                         uDisplayInformation: LongWord) of object;
  TMsRdpClient5NotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                        windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient5NotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient5NotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient5NotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient5NotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient5NotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient5NotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                       bandwidth: Integer; 
                                                                                       rtt: Integer) of object;
  TMsRdpClient5NotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                    networkAvailable: WordBool; 
                                                                                    attemptCount: Integer; 
                                                                                    maxAttemptCount: Integer) of object;

  TMsRdpClient5NotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient5NotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient5NotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient5NotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClient5NotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient5NotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient5NotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient5NotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient5NotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient5NotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient5NotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient5NotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClient5NotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient5NotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient5NotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient5NotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient5NotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient5NotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient5;
    function  GetControlInterface: IMsRdpClient5;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    property  ControlInterface: IMsRdpClient5 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient5 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient5NotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient5NotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient5NotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient5NotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient5NotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient5NotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient5NotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient5NotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient5NotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient5NotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient5NotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient5NotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient5NotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient5NotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient5NotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient5NotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient5NotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient5
// Help String      : Microsoft RDP Client Control (redistributable) - version 6
// Default Interface: IMsRdpClient5
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient5OnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient5OnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                   const data: WideString) of object;
  TMsRdpClient5OnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient5OnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient5OnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                       height: Integer) of object;
  TMsRdpClient5OnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient5OnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                attemptCount: Integer) of object;
  TMsRdpClient5OnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                   lError: RemoteProgramResult; 
                                                                   vbIsExecutable: WordBool) of object;
  TMsRdpClient5OnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      uDisplayInformation: LongWord) of object;
  TMsRdpClient5OnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                     windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient5OnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient5OnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient5OnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient5OnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient5OnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient5OnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                    bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClient5OnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 networkAvailable: WordBool; 
                                                                 attemptCount: Integer; 
                                                                 maxAttemptCount: Integer) of object;

  TMsRdpClient5 = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient5OnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient5OnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient5OnFatalError;
    FOnWarning: TMsRdpClient5OnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient5OnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient5OnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient5OnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient5OnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient5OnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient5OnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient5OnLogonError;
    FOnFocusReleased: TMsRdpClient5OnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient5OnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient5OnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient5OnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient5OnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient5OnAutoReconnecting2;
    FIntf: IMsRdpClient5;
    function  GetControlInterface: IMsRdpClient5;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    property  ControlInterface: IMsRdpClient5 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient5 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient5OnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient5OnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient5OnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient5OnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient5OnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient5OnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient5OnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient5OnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient5OnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient5OnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient5OnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient5OnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient5OnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient5OnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient5OnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient5OnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient5OnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient6NotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 7
// Default Interface: IMsRdpClient6
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient6NotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient6NotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                      const data: WideString) of object;
  TMsRdpClient6NotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient6NotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient6NotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                          height: Integer) of object;
  TMsRdpClient6NotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient6NotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   attemptCount: Integer) of object;
  TMsRdpClient6NotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                      lError: RemoteProgramResult; 
                                                                                      vbIsExecutable: WordBool) of object;
  TMsRdpClient6NotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                         uDisplayInformation: LongWord) of object;
  TMsRdpClient6NotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                        windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient6NotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient6NotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient6NotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient6NotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient6NotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient6NotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                       bandwidth: Integer; 
                                                                                       rtt: Integer) of object;
  TMsRdpClient6NotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                    networkAvailable: WordBool; 
                                                                                    attemptCount: Integer; 
                                                                                    maxAttemptCount: Integer) of object;

  TMsRdpClient6NotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient6NotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient6NotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient6NotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClient6NotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient6NotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient6NotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient6NotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient6NotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient6NotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient6NotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient6NotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClient6NotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient6NotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient6NotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient6NotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient6NotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient6NotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient6;
    function  GetControlInterface: IMsRdpClient6;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    property  ControlInterface: IMsRdpClient6 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient6 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient6NotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient6NotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient6NotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient6NotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient6NotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient6NotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient6NotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient6NotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient6NotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient6NotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient6NotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient6NotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient6NotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient6NotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient6NotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient6NotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient6NotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient6
// Help String      : Microsoft RDP Client Control (redistributable) - version 7
// Default Interface: IMsRdpClient6
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient6OnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient6OnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                   const data: WideString) of object;
  TMsRdpClient6OnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient6OnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient6OnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                       height: Integer) of object;
  TMsRdpClient6OnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient6OnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                attemptCount: Integer) of object;
  TMsRdpClient6OnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                   lError: RemoteProgramResult; 
                                                                   vbIsExecutable: WordBool) of object;
  TMsRdpClient6OnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      uDisplayInformation: LongWord) of object;
  TMsRdpClient6OnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                     windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient6OnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient6OnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient6OnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient6OnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient6OnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient6OnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                    bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClient6OnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 networkAvailable: WordBool; 
                                                                 attemptCount: Integer; 
                                                                 maxAttemptCount: Integer) of object;

  TMsRdpClient6 = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient6OnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient6OnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient6OnFatalError;
    FOnWarning: TMsRdpClient6OnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient6OnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient6OnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient6OnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient6OnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient6OnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient6OnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient6OnLogonError;
    FOnFocusReleased: TMsRdpClient6OnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient6OnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient6OnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient6OnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient6OnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient6OnAutoReconnecting2;
    FIntf: IMsRdpClient6;
    function  GetControlInterface: IMsRdpClient6;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    property  ControlInterface: IMsRdpClient6 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient6 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient6OnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient6OnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient6OnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient6OnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient6OnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient6OnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient6OnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient6OnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient6OnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient6OnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient6OnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient6OnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient6OnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient6OnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient6OnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient6OnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient6OnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient7NotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 8
// Default Interface: IMsRdpClient7
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (34) CanCreate Control
// *********************************************************************//
  TMsRdpClient7NotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient7NotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                      const data: WideString) of object;
  TMsRdpClient7NotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient7NotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient7NotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                          height: Integer) of object;
  TMsRdpClient7NotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient7NotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   attemptCount: Integer) of object;
  TMsRdpClient7NotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                      lError: RemoteProgramResult; 
                                                                                      vbIsExecutable: WordBool) of object;
  TMsRdpClient7NotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                         uDisplayInformation: LongWord) of object;
  TMsRdpClient7NotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                        windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient7NotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient7NotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient7NotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient7NotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient7NotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient7NotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                       bandwidth: Integer; 
                                                                                       rtt: Integer) of object;
  TMsRdpClient7NotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                    networkAvailable: WordBool; 
                                                                                    attemptCount: Integer; 
                                                                                    maxAttemptCount: Integer) of object;

  TMsRdpClient7NotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient7NotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient7NotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient7NotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClient7NotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient7NotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient7NotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient7NotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient7NotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient7NotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient7NotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient7NotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClient7NotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient7NotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient7NotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient7NotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient7NotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient7NotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient7;
    function  GetControlInterface: IMsRdpClient7;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2;
    function Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
    function Get_TransportSettings3: IMsRdpClientTransportSettings3;
    function Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
    function Get_RemoteProgram2: ITSRemoteProgram2;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    function GetStatusText(statusCode: SYSUINT): WideString;
    property  ControlInterface: IMsRdpClient7 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient7 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 read Get_AdvancedSettings8;
    property TransportSettings3: IMsRdpClientTransportSettings3 read Get_TransportSettings3;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 read Get_SecuredSettings3;
    property RemoteProgram2: ITSRemoteProgram2 read Get_RemoteProgram2;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient7NotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient7NotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient7NotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient7NotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient7NotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient7NotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient7NotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient7NotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient7NotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient7NotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient7NotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient7NotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient7NotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient7NotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient7NotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient7NotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient7NotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient7
// Help String      : Microsoft RDP Client Control (redistributable) - version 8
// Default Interface: IMsRdpClient7
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (2) CanCreate
// *********************************************************************//
  TMsRdpClient7OnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient7OnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                   const data: WideString) of object;
  TMsRdpClient7OnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient7OnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient7OnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                       height: Integer) of object;
  TMsRdpClient7OnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient7OnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                attemptCount: Integer) of object;
  TMsRdpClient7OnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                   lError: RemoteProgramResult; 
                                                                   vbIsExecutable: WordBool) of object;
  TMsRdpClient7OnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      uDisplayInformation: LongWord) of object;
  TMsRdpClient7OnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                     windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient7OnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient7OnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient7OnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient7OnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient7OnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient7OnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                    bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClient7OnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 networkAvailable: WordBool; 
                                                                 attemptCount: Integer; 
                                                                 maxAttemptCount: Integer) of object;

  TMsRdpClient7 = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient7OnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient7OnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient7OnFatalError;
    FOnWarning: TMsRdpClient7OnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient7OnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient7OnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient7OnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient7OnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient7OnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient7OnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient7OnLogonError;
    FOnFocusReleased: TMsRdpClient7OnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient7OnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient7OnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient7OnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient7OnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient7OnAutoReconnecting2;
    FIntf: IMsRdpClient7;
    function  GetControlInterface: IMsRdpClient7;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2;
    function Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
    function Get_TransportSettings3: IMsRdpClientTransportSettings3;
    function Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
    function Get_RemoteProgram2: ITSRemoteProgram2;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    function GetStatusText(statusCode: SYSUINT): WideString;
    property  ControlInterface: IMsRdpClient7 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient7 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 read Get_AdvancedSettings8;
    property TransportSettings3: IMsRdpClientTransportSettings3 read Get_TransportSettings3;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 read Get_SecuredSettings3;
    property RemoteProgram2: ITSRemoteProgram2 read Get_RemoteProgram2;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient7OnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient7OnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient7OnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient7OnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient7OnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient7OnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient7OnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient7OnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient7OnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient7OnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient7OnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient7OnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient7OnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient7OnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient7OnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient7OnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient7OnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient8NotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 9
// Default Interface: IMsRdpClient8
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (2) CanCreate
// *********************************************************************//
  TMsRdpClient8NotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient8NotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                      const data: WideString) of object;
  TMsRdpClient8NotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient8NotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient8NotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                          height: Integer) of object;
  TMsRdpClient8NotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient8NotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   attemptCount: Integer) of object;
  TMsRdpClient8NotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                      lError: RemoteProgramResult; 
                                                                                      vbIsExecutable: WordBool) of object;
  TMsRdpClient8NotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                         uDisplayInformation: LongWord) of object;
  TMsRdpClient8NotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                        windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient8NotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient8NotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient8NotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient8NotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient8NotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient8NotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                       bandwidth: Integer; 
                                                                                       rtt: Integer) of object;
  TMsRdpClient8NotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                    networkAvailable: WordBool; 
                                                                                    attemptCount: Integer; 
                                                                                    maxAttemptCount: Integer) of object;

  TMsRdpClient8NotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient8NotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient8NotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient8NotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClient8NotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient8NotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient8NotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient8NotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient8NotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient8NotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient8NotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient8NotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClient8NotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient8NotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient8NotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient8NotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient8NotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient8NotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient8;
    function  GetControlInterface: IMsRdpClient8;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2;
    function Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
    function Get_TransportSettings3: IMsRdpClientTransportSettings3;
    function Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
    function Get_RemoteProgram2: ITSRemoteProgram2;
    function Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    function GetStatusText(statusCode: SYSUINT): WideString;
    procedure SendRemoteAction(actionType: RemoteSessionActionType);
    function Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus;
    property  ControlInterface: IMsRdpClient8 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient8 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 read Get_AdvancedSettings8;
    property TransportSettings3: IMsRdpClientTransportSettings3 read Get_TransportSettings3;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 read Get_SecuredSettings3;
    property RemoteProgram2: ITSRemoteProgram2 read Get_RemoteProgram2;
    property AdvancedSettings9: IMsRdpClientAdvancedSettings8 read Get_AdvancedSettings9;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient8NotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient8NotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient8NotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient8NotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient8NotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient8NotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient8NotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient8NotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient8NotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient8NotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient8NotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient8NotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient8NotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient8NotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient8NotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient8NotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient8NotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient8
// Help String      : Microsoft RDP Client Control (redistributable) - version 9
// Default Interface: IMsRdpClient8
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (2) CanCreate
// *********************************************************************//
  TMsRdpClient8OnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient8OnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                   const data: WideString) of object;
  TMsRdpClient8OnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient8OnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient8OnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                       height: Integer) of object;
  TMsRdpClient8OnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient8OnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                attemptCount: Integer) of object;
  TMsRdpClient8OnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                   lError: RemoteProgramResult; 
                                                                   vbIsExecutable: WordBool) of object;
  TMsRdpClient8OnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      uDisplayInformation: LongWord) of object;
  TMsRdpClient8OnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                     windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient8OnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient8OnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient8OnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient8OnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient8OnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient8OnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                    bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClient8OnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 networkAvailable: WordBool; 
                                                                 attemptCount: Integer; 
                                                                 maxAttemptCount: Integer) of object;

  TMsRdpClient8 = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient8OnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient8OnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient8OnFatalError;
    FOnWarning: TMsRdpClient8OnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient8OnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient8OnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient8OnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient8OnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient8OnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient8OnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient8OnLogonError;
    FOnFocusReleased: TMsRdpClient8OnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient8OnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient8OnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient8OnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient8OnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient8OnAutoReconnecting2;
    FIntf: IMsRdpClient8;
    function  GetControlInterface: IMsRdpClient8;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2;
    function Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
    function Get_TransportSettings3: IMsRdpClientTransportSettings3;
    function Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
    function Get_RemoteProgram2: ITSRemoteProgram2;
    function Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    function GetStatusText(statusCode: SYSUINT): WideString;
    procedure SendRemoteAction(actionType: RemoteSessionActionType);
    function Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus;
    property  ControlInterface: IMsRdpClient8 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient8 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 read Get_AdvancedSettings8;
    property TransportSettings3: IMsRdpClientTransportSettings3 read Get_TransportSettings3;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 read Get_SecuredSettings3;
    property RemoteProgram2: ITSRemoteProgram2 read Get_RemoteProgram2;
    property AdvancedSettings9: IMsRdpClientAdvancedSettings8 read Get_AdvancedSettings9;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient8OnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient8OnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient8OnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient8OnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient8OnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient8OnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient8OnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient8OnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient8OnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient8OnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient8OnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient8OnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient8OnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient8OnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient8OnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient8OnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient8OnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient9NotSafeForScripting
// Help String      : Microsoft RDP Client Control - version 10
// Default Interface: IMsRdpClient9
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (2) CanCreate
// *********************************************************************//
  TMsRdpClient9NotSafeForScriptingOnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient9NotSafeForScriptingOnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                                      const data: WideString) of object;
  TMsRdpClient9NotSafeForScriptingOnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient9NotSafeForScriptingOnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient9NotSafeForScriptingOnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                                          height: Integer) of object;
  TMsRdpClient9NotSafeForScriptingOnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient9NotSafeForScriptingOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                   attemptCount: Integer) of object;
  TMsRdpClient9NotSafeForScriptingOnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                                      lError: RemoteProgramResult; 
                                                                                      vbIsExecutable: WordBool) of object;
  TMsRdpClient9NotSafeForScriptingOnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                         uDisplayInformation: LongWord) of object;
  TMsRdpClient9NotSafeForScriptingOnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                                        var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                                        windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient9NotSafeForScriptingOnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient9NotSafeForScriptingOnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient9NotSafeForScriptingOnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient9NotSafeForScriptingOnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient9NotSafeForScriptingOnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient9NotSafeForScriptingOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                                       bandwidth: Integer; 
                                                                                       rtt: Integer) of object;
  TMsRdpClient9NotSafeForScriptingOnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                                    networkAvailable: WordBool; 
                                                                                    attemptCount: Integer; 
                                                                                    maxAttemptCount: Integer) of object;

  TMsRdpClient9NotSafeForScripting = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient9NotSafeForScriptingOnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient9NotSafeForScriptingOnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient9NotSafeForScriptingOnFatalError;
    FOnWarning: TMsRdpClient9NotSafeForScriptingOnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient9NotSafeForScriptingOnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient9NotSafeForScriptingOnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient9NotSafeForScriptingOnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient9NotSafeForScriptingOnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient9NotSafeForScriptingOnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient9NotSafeForScriptingOnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient9NotSafeForScriptingOnLogonError;
    FOnFocusReleased: TMsRdpClient9NotSafeForScriptingOnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient9NotSafeForScriptingOnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient9NotSafeForScriptingOnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient9NotSafeForScriptingOnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient9NotSafeForScriptingOnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient9NotSafeForScriptingOnAutoReconnecting2;
    FIntf: IMsRdpClient9;
    function  GetControlInterface: IMsRdpClient9;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2;
    function Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
    function Get_TransportSettings3: IMsRdpClientTransportSettings3;
    function Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
    function Get_RemoteProgram2: ITSRemoteProgram2;
    function Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8;
    function Get_TransportSettings4: IMsRdpClientTransportSettings4;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    function GetStatusText(statusCode: SYSUINT): WideString;
    procedure SendRemoteAction(actionType: RemoteSessionActionType);
    function Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus;
    procedure SyncSessionDisplaySettings;
    procedure UpdateSessionDisplaySettings(ulDesktopWidth: LongWord; ulDesktopHeight: LongWord; 
                                           ulPhysicalWidth: LongWord; ulPhysicalHeight: LongWord; 
                                           ulOrientation: LongWord; ulDesktopScaleFactor: LongWord; 
                                           ulDeviceScaleFactor: LongWord);
    procedure attachEvent(const eventName: WideString; const callback: IDispatch);
    procedure detachEvent(const eventName: WideString; const callback: IDispatch);
    property  ControlInterface: IMsRdpClient9 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient9 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 read Get_AdvancedSettings8;
    property TransportSettings3: IMsRdpClientTransportSettings3 read Get_TransportSettings3;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 read Get_SecuredSettings3;
    property RemoteProgram2: ITSRemoteProgram2 read Get_RemoteProgram2;
    property AdvancedSettings9: IMsRdpClientAdvancedSettings8 read Get_AdvancedSettings9;
    property TransportSettings4: IMsRdpClientTransportSettings4 read Get_TransportSettings4;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient9NotSafeForScriptingOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient9NotSafeForScriptingOnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient9NotSafeForScriptingOnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient9NotSafeForScriptingOnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient9NotSafeForScriptingOnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient9NotSafeForScriptingOnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient9NotSafeForScriptingOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient9NotSafeForScriptingOnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient9NotSafeForScriptingOnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient9NotSafeForScriptingOnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient9NotSafeForScriptingOnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient9NotSafeForScriptingOnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient9NotSafeForScriptingOnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient9NotSafeForScriptingOnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient9NotSafeForScriptingOnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient9NotSafeForScriptingOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient9NotSafeForScriptingOnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;


// *********************************************************************//
// OLE Control Proxy class declaration
// Control Name     : TMsRdpClient9
// Help String      : Microsoft RDP Client Control (redistributable) - version 10
// Default Interface: IMsRdpClient9
// Def. Intf. DISP? : No
// Event   Interface: IMsTscAxEvents
// TypeFlags        : (2) CanCreate
// *********************************************************************//
  TMsRdpClient9OnDisconnected = procedure(ASender: TObject; discReason: Integer) of object;
  TMsRdpClient9OnChannelReceivedData = procedure(ASender: TObject; const chanName: WideString; 
                                                                   const data: WideString) of object;
  TMsRdpClient9OnFatalError = procedure(ASender: TObject; errorCode: Integer) of object;
  TMsRdpClient9OnWarning = procedure(ASender: TObject; warningCode: Integer) of object;
  TMsRdpClient9OnRemoteDesktopSizeChange = procedure(ASender: TObject; width: Integer; 
                                                                       height: Integer) of object;
  TMsRdpClient9OnReceivedTSPublicKey = procedure(ASender: TObject; const publicKey: WideString) of object;
  TMsRdpClient9OnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                attemptCount: Integer) of object;
  TMsRdpClient9OnRemoteProgramResult = procedure(ASender: TObject; const bstrRemoteProgram: WideString; 
                                                                   lError: RemoteProgramResult; 
                                                                   vbIsExecutable: WordBool) of object;
  TMsRdpClient9OnRemoteProgramDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                      uDisplayInformation: LongWord) of object;
  TMsRdpClient9OnRemoteWindowDisplayed = procedure(ASender: TObject; vbDisplayed: WordBool; 
                                                                     var hwnd: {NOT_OLEAUTO(_RemotableHandle)}OleVariant; 
                                                                     windowAttribute: RemoteWindowDisplayedAttribute) of object;
  TMsRdpClient9OnLogonError = procedure(ASender: TObject; lError: Integer) of object;
  TMsRdpClient9OnFocusReleased = procedure(ASender: TObject; iDirection: SYSINT) of object;
  TMsRdpClient9OnUserNameAcquired = procedure(ASender: TObject; const bstrUserName: WideString) of object;
  TMsRdpClient9OnMouseInputModeChanged = procedure(ASender: TObject; fMouseModeRelative: WordBool) of object;
  TMsRdpClient9OnServiceMessageReceived = procedure(ASender: TObject; const serviceMessage: WideString) of object;
  TMsRdpClient9OnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                    bandwidth: Integer; rtt: Integer) of object;
  TMsRdpClient9OnAutoReconnecting2 = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                 networkAvailable: WordBool; 
                                                                 attemptCount: Integer; 
                                                                 maxAttemptCount: Integer) of object;

  TMsRdpClient9 = class(TOleControl)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginComplete: TNotifyEvent;
    FOnDisconnected: TMsRdpClient9OnDisconnected;
    FOnEnterFullScreenMode: TNotifyEvent;
    FOnLeaveFullScreenMode: TNotifyEvent;
    FOnChannelReceivedData: TMsRdpClient9OnChannelReceivedData;
    FOnRequestGoFullScreen: TNotifyEvent;
    FOnRequestLeaveFullScreen: TNotifyEvent;
    FOnFatalError: TMsRdpClient9OnFatalError;
    FOnWarning: TMsRdpClient9OnWarning;
    FOnRemoteDesktopSizeChange: TMsRdpClient9OnRemoteDesktopSizeChange;
    FOnIdleTimeoutNotification: TNotifyEvent;
    FOnRequestContainerMinimize: TNotifyEvent;
    FOnConfirmClose: TNotifyEvent;
    FOnReceivedTSPublicKey: TMsRdpClient9OnReceivedTSPublicKey;
    FOnAutoReconnecting: TMsRdpClient9OnAutoReconnecting;
    FOnAuthenticationWarningDisplayed: TNotifyEvent;
    FOnAuthenticationWarningDismissed: TNotifyEvent;
    FOnRemoteProgramResult: TMsRdpClient9OnRemoteProgramResult;
    FOnRemoteProgramDisplayed: TMsRdpClient9OnRemoteProgramDisplayed;
    FOnRemoteWindowDisplayed: TMsRdpClient9OnRemoteWindowDisplayed;
    FOnLogonError: TMsRdpClient9OnLogonError;
    FOnFocusReleased: TMsRdpClient9OnFocusReleased;
    FOnUserNameAcquired: TMsRdpClient9OnUserNameAcquired;
    FOnMouseInputModeChanged: TMsRdpClient9OnMouseInputModeChanged;
    FOnServiceMessageReceived: TMsRdpClient9OnServiceMessageReceived;
    FOnConnectionBarPullDown: TNotifyEvent;
    FOnNetworkStatusChanged: TMsRdpClient9OnNetworkStatusChanged;
    FOnDevicesButtonPressed: TNotifyEvent;
    FOnAutoReconnected: TNotifyEvent;
    FOnAutoReconnecting2: TMsRdpClient9OnAutoReconnecting2;
    FIntf: IMsRdpClient9;
    function  GetControlInterface: IMsRdpClient9;
  protected
    procedure CreateControl;
    procedure InitControlData; override;
    function Get_SecuredSettings: IMsTscSecuredSettings;
    function Get_AdvancedSettings: IMsTscAdvancedSettings;
    function Get_Debugger: IMsTscDebug;
    function Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
    function Get_SecuredSettings2: IMsRdpClientSecuredSettings;
    function Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
    function Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
    function Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
    function Get_TransportSettings: IMsRdpClientTransportSettings;
    function Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
    function Get_RemoteProgram: ITSRemoteProgram;
    function Get_MsRdpClientShell: IMsRdpClientShell;
    function Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
    function Get_TransportSettings2: IMsRdpClientTransportSettings2;
    function Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
    function Get_TransportSettings3: IMsRdpClientTransportSettings3;
    function Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
    function Get_RemoteProgram2: ITSRemoteProgram2;
    function Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8;
    function Get_TransportSettings4: IMsRdpClientTransportSettings4;
  public
    procedure Connect;
    procedure Disconnect;
    procedure CreateVirtualChannels(const newVal: WideString);
    procedure SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
    procedure SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
    function GetVirtualChannelOptions(const chanName: WideString): Integer;
    function RequestClose: ControlCloseStatus;
    function GetErrorDescription(disconnectReason: SYSUINT; ExtendedDisconnectReason: SYSUINT): WideString;
    function GetStatusText(statusCode: SYSUINT): WideString;
    procedure SendRemoteAction(actionType: RemoteSessionActionType);
    function Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus;
    procedure SyncSessionDisplaySettings;
    procedure UpdateSessionDisplaySettings(ulDesktopWidth: LongWord; ulDesktopHeight: LongWord; 
                                           ulPhysicalWidth: LongWord; ulPhysicalHeight: LongWord; 
                                           ulOrientation: LongWord; ulDesktopScaleFactor: LongWord; 
                                           ulDeviceScaleFactor: LongWord);
    procedure attachEvent(const eventName: WideString; const callback: IDispatch);
    procedure detachEvent(const eventName: WideString; const callback: IDispatch);
    property  ControlInterface: IMsRdpClient9 read GetControlInterface;
    property  DefaultInterface: IMsRdpClient9 read GetControlInterface;
    property Connected: Smallint index 6 read GetSmallintProp;
    property HorizontalScrollBarVisible: Integer index 17 read GetIntegerProp;
    property VerticalScrollBarVisible: Integer index 18 read GetIntegerProp;
    property FullScreenTitle: WideString index 19 write SetWideStringProp;
    property CipherStrength: Integer index 20 read GetIntegerProp;
    property Version: WideString index 21 read GetWideStringProp;
    property SecuredSettingsEnabled: Integer index 22 read GetIntegerProp;
    property SecuredSettings: IMsTscSecuredSettings read Get_SecuredSettings;
    property AdvancedSettings: IMsTscAdvancedSettings read Get_AdvancedSettings;
    property Debugger: IMsTscDebug read Get_Debugger;
    property AdvancedSettings2: IMsRdpClientAdvancedSettings read Get_AdvancedSettings2;
    property SecuredSettings2: IMsRdpClientSecuredSettings read Get_SecuredSettings2;
    property ExtendedDisconnectReason: TOleEnum index 103 read GetTOleEnumProp;
    property AdvancedSettings3: IMsRdpClientAdvancedSettings2 read Get_AdvancedSettings3;
    property AdvancedSettings4: IMsRdpClientAdvancedSettings3 read Get_AdvancedSettings4;
    property AdvancedSettings5: IMsRdpClientAdvancedSettings4 read Get_AdvancedSettings5;
    property TransportSettings: IMsRdpClientTransportSettings read Get_TransportSettings;
    property AdvancedSettings6: IMsRdpClientAdvancedSettings5 read Get_AdvancedSettings6;
    property RemoteProgram: ITSRemoteProgram read Get_RemoteProgram;
    property MsRdpClientShell: IMsRdpClientShell read Get_MsRdpClientShell;
    property AdvancedSettings7: IMsRdpClientAdvancedSettings6 read Get_AdvancedSettings7;
    property TransportSettings2: IMsRdpClientTransportSettings2 read Get_TransportSettings2;
    property AdvancedSettings8: IMsRdpClientAdvancedSettings7 read Get_AdvancedSettings8;
    property TransportSettings3: IMsRdpClientTransportSettings3 read Get_TransportSettings3;
    property SecuredSettings3: IMsRdpClientSecuredSettings2 read Get_SecuredSettings3;
    property RemoteProgram2: ITSRemoteProgram2 read Get_RemoteProgram2;
    property AdvancedSettings9: IMsRdpClientAdvancedSettings8 read Get_AdvancedSettings9;
    property TransportSettings4: IMsRdpClientTransportSettings4 read Get_TransportSettings4;
  published
    property Anchors;
    property  TabStop;
    property  Align;
    property  DragCursor;
    property  DragMode;
    property  ParentShowHint;
    property  PopupMenu;
    property  ShowHint;
    property  TabOrder;
    property  Visible;
    property  OnDragDrop;
    property  OnDragOver;
    property  OnEndDrag;
    property  OnEnter;
    property  OnExit;
    property  OnStartDrag;
    property Server: WideString index 1 read GetWideStringProp write SetWideStringProp stored False;
    property Domain: WideString index 2 read GetWideStringProp write SetWideStringProp stored False;
    property UserName: WideString index 3 read GetWideStringProp write SetWideStringProp stored False;
    property DisconnectedText: WideString index 4 read GetWideStringProp write SetWideStringProp stored False;
    property ConnectingText: WideString index 5 read GetWideStringProp write SetWideStringProp stored False;
    property DesktopWidth: Integer index 12 read GetIntegerProp write SetIntegerProp stored False;
    property DesktopHeight: Integer index 13 read GetIntegerProp write SetIntegerProp stored False;
    property StartConnected: Integer index 16 read GetIntegerProp write SetIntegerProp stored False;
    property ColorDepth: Integer index 100 read GetIntegerProp write SetIntegerProp stored False;
    property FullScreen: WordBool index 104 read GetWordBoolProp write SetWordBoolProp stored False;
    property ConnectedStatusText: WideString index 201 read GetWideStringProp write SetWideStringProp stored False;
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginComplete: TNotifyEvent read FOnLoginComplete write FOnLoginComplete;
    property OnDisconnected: TMsRdpClient9OnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnEnterFullScreenMode: TNotifyEvent read FOnEnterFullScreenMode write FOnEnterFullScreenMode;
    property OnLeaveFullScreenMode: TNotifyEvent read FOnLeaveFullScreenMode write FOnLeaveFullScreenMode;
    property OnChannelReceivedData: TMsRdpClient9OnChannelReceivedData read FOnChannelReceivedData write FOnChannelReceivedData;
    property OnRequestGoFullScreen: TNotifyEvent read FOnRequestGoFullScreen write FOnRequestGoFullScreen;
    property OnRequestLeaveFullScreen: TNotifyEvent read FOnRequestLeaveFullScreen write FOnRequestLeaveFullScreen;
    property OnFatalError: TMsRdpClient9OnFatalError read FOnFatalError write FOnFatalError;
    property OnWarning: TMsRdpClient9OnWarning read FOnWarning write FOnWarning;
    property OnRemoteDesktopSizeChange: TMsRdpClient9OnRemoteDesktopSizeChange read FOnRemoteDesktopSizeChange write FOnRemoteDesktopSizeChange;
    property OnIdleTimeoutNotification: TNotifyEvent read FOnIdleTimeoutNotification write FOnIdleTimeoutNotification;
    property OnRequestContainerMinimize: TNotifyEvent read FOnRequestContainerMinimize write FOnRequestContainerMinimize;
    property OnConfirmClose: TNotifyEvent read FOnConfirmClose write FOnConfirmClose;
    property OnReceivedTSPublicKey: TMsRdpClient9OnReceivedTSPublicKey read FOnReceivedTSPublicKey write FOnReceivedTSPublicKey;
    property OnAutoReconnecting: TMsRdpClient9OnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAuthenticationWarningDisplayed: TNotifyEvent read FOnAuthenticationWarningDisplayed write FOnAuthenticationWarningDisplayed;
    property OnAuthenticationWarningDismissed: TNotifyEvent read FOnAuthenticationWarningDismissed write FOnAuthenticationWarningDismissed;
    property OnRemoteProgramResult: TMsRdpClient9OnRemoteProgramResult read FOnRemoteProgramResult write FOnRemoteProgramResult;
    property OnRemoteProgramDisplayed: TMsRdpClient9OnRemoteProgramDisplayed read FOnRemoteProgramDisplayed write FOnRemoteProgramDisplayed;
    property OnRemoteWindowDisplayed: TMsRdpClient9OnRemoteWindowDisplayed read FOnRemoteWindowDisplayed write FOnRemoteWindowDisplayed;
    property OnLogonError: TMsRdpClient9OnLogonError read FOnLogonError write FOnLogonError;
    property OnFocusReleased: TMsRdpClient9OnFocusReleased read FOnFocusReleased write FOnFocusReleased;
    property OnUserNameAcquired: TMsRdpClient9OnUserNameAcquired read FOnUserNameAcquired write FOnUserNameAcquired;
    property OnMouseInputModeChanged: TMsRdpClient9OnMouseInputModeChanged read FOnMouseInputModeChanged write FOnMouseInputModeChanged;
    property OnServiceMessageReceived: TMsRdpClient9OnServiceMessageReceived read FOnServiceMessageReceived write FOnServiceMessageReceived;
    property OnConnectionBarPullDown: TNotifyEvent read FOnConnectionBarPullDown write FOnConnectionBarPullDown;
    property OnNetworkStatusChanged: TMsRdpClient9OnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnDevicesButtonPressed: TNotifyEvent read FOnDevicesButtonPressed write FOnDevicesButtonPressed;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnAutoReconnecting2: TMsRdpClient9OnAutoReconnecting2 read FOnAutoReconnecting2 write FOnAutoReconnecting2;
  end;

// *********************************************************************//
// The Class CoRemoteDesktopClient provides a Create and CreateRemote method to          
// create instances of the default interface IRemoteDesktopClient exposed by              
// the CoClass RemoteDesktopClient. The functions are intended to be used by             
// clients wishing to automate the CoClass objects exposed by the         
// server of this typelibrary.                                            
// *********************************************************************//
  CoRemoteDesktopClient = class
    class function Create: IRemoteDesktopClient;
    class function CreateRemote(const MachineName: string): IRemoteDesktopClient;
  end;

  TRemoteDesktopClientOnDisconnected = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                   ExtendedDisconnectReason: Integer; 
                                                                   const disconnectErrorMessage: WideString) of object;
  TRemoteDesktopClientOnStatusChanged = procedure(ASender: TObject; statusCode: Integer; 
                                                                    const statusMessage: WideString) of object;
  TRemoteDesktopClientOnAutoReconnecting = procedure(ASender: TObject; disconnectReason: Integer; 
                                                                       ExtendedDisconnectReason: Integer; 
                                                                       const disconnectErrorMessage: WideString; 
                                                                       networkAvailable: WordBool; 
                                                                       attemptCount: Integer; 
                                                                       maxAttemptCount: Integer) of object;
  TRemoteDesktopClientOnNetworkStatusChanged = procedure(ASender: TObject; qualityLevel: LongWord; 
                                                                           bandwidth: Integer; 
                                                                           rtt: Integer) of object;
  TRemoteDesktopClientOnAdminMessageReceived = procedure(ASender: TObject; const adminMessage: WideString) of object;
  TRemoteDesktopClientOnKeyCombinationPressed = procedure(ASender: TObject; keyCombination: Integer) of object;
  TRemoteDesktopClientOnRemoteDesktopSizeChanged = procedure(ASender: TObject; width: Integer; 
                                                                               height: Integer) of object;
  TRemoteDesktopClientOnTouchPointerCursorMoved = procedure(ASender: TObject; x: Integer; y: Integer) of object;


// *********************************************************************//
// OLE Server Proxy class declaration
// Server Object    : TRemoteDesktopClient
// Help String      : Microsoft Metro Style Remote Desktop Client Control - version 1
// Default Interface: IRemoteDesktopClient
// Def. Intf. DISP? : No
// Event   Interface: IRemoteDesktopClientEvents
// TypeFlags        : (2) CanCreate
// *********************************************************************//
  TRemoteDesktopClient = class(TOleServer)
  private
    FOnConnecting: TNotifyEvent;
    FOnConnected: TNotifyEvent;
    FOnLoginCompleted: TNotifyEvent;
    FOnDisconnected: TRemoteDesktopClientOnDisconnected;
    FOnStatusChanged: TRemoteDesktopClientOnStatusChanged;
    FOnAutoReconnecting: TRemoteDesktopClientOnAutoReconnecting;
    FOnAutoReconnected: TNotifyEvent;
    FOnDialogDisplaying: TNotifyEvent;
    FOnDialogDismissed: TNotifyEvent;
    FOnNetworkStatusChanged: TRemoteDesktopClientOnNetworkStatusChanged;
    FOnAdminMessageReceived: TRemoteDesktopClientOnAdminMessageReceived;
    FOnKeyCombinationPressed: TRemoteDesktopClientOnKeyCombinationPressed;
    FOnRemoteDesktopSizeChanged: TRemoteDesktopClientOnRemoteDesktopSizeChanged;
    FOnTouchPointerCursorMoved: TRemoteDesktopClientOnTouchPointerCursorMoved;
    FIntf: IRemoteDesktopClient;
    function GetDefaultInterface: IRemoteDesktopClient;
  protected
    procedure InitServerData; override;
    procedure InvokeEvent(DispID: TDispID; var Params: TVariantArray); override;
    function Get_Settings: IRemoteDesktopClientSettings;
    function Get_Actions: IRemoteDesktopClientActions;
    function Get_TouchPointer: IRemoteDesktopClientTouchPointer;
  public
    constructor Create(AOwner: TComponent); override;
    destructor  Destroy; override;
    procedure Connect; override;
    procedure ConnectTo(svrIntf: IRemoteDesktopClient);
    procedure Disconnect; override;
    procedure Connect1;
    procedure Disconnect1;
    procedure Reconnect(width: LongWord; height: LongWord);
    procedure DeleteSavedCredentials(const serverName: WideString);
    procedure UpdateSessionDisplaySettings(width: LongWord; height: LongWord);
    procedure attachEvent(const eventName: WideString; const callback: IDispatch);
    procedure detachEvent(const eventName: WideString; const callback: IDispatch);
    property DefaultInterface: IRemoteDesktopClient read GetDefaultInterface;
    property Settings: IRemoteDesktopClientSettings read Get_Settings;
    property Actions: IRemoteDesktopClientActions read Get_Actions;
    property TouchPointer: IRemoteDesktopClientTouchPointer read Get_TouchPointer;
  published
    property OnConnecting: TNotifyEvent read FOnConnecting write FOnConnecting;
    property OnConnected: TNotifyEvent read FOnConnected write FOnConnected;
    property OnLoginCompleted: TNotifyEvent read FOnLoginCompleted write FOnLoginCompleted;
    property OnDisconnected: TRemoteDesktopClientOnDisconnected read FOnDisconnected write FOnDisconnected;
    property OnStatusChanged: TRemoteDesktopClientOnStatusChanged read FOnStatusChanged write FOnStatusChanged;
    property OnAutoReconnecting: TRemoteDesktopClientOnAutoReconnecting read FOnAutoReconnecting write FOnAutoReconnecting;
    property OnAutoReconnected: TNotifyEvent read FOnAutoReconnected write FOnAutoReconnected;
    property OnDialogDisplaying: TNotifyEvent read FOnDialogDisplaying write FOnDialogDisplaying;
    property OnDialogDismissed: TNotifyEvent read FOnDialogDismissed write FOnDialogDismissed;
    property OnNetworkStatusChanged: TRemoteDesktopClientOnNetworkStatusChanged read FOnNetworkStatusChanged write FOnNetworkStatusChanged;
    property OnAdminMessageReceived: TRemoteDesktopClientOnAdminMessageReceived read FOnAdminMessageReceived write FOnAdminMessageReceived;
    property OnKeyCombinationPressed: TRemoteDesktopClientOnKeyCombinationPressed read FOnKeyCombinationPressed write FOnKeyCombinationPressed;
    property OnRemoteDesktopSizeChanged: TRemoteDesktopClientOnRemoteDesktopSizeChanged read FOnRemoteDesktopSizeChanged write FOnRemoteDesktopSizeChanged;
    property OnTouchPointerCursorMoved: TRemoteDesktopClientOnTouchPointerCursorMoved read FOnTouchPointerCursorMoved write FOnTouchPointerCursorMoved;
  end;

procedure Register;

resourcestring
  dtlServerPage = 'ActiveX';

  dtlOcxPage = 'ActiveX';

implementation

uses System.Win.ComObj;

procedure TMsTscAxNotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{A41A4187-5A86-4E26-B40A-856F9035D9CB}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsTscAxNotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsTscAx;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsTscAxNotSafeForScripting.GetControlInterface: IMsTscAx;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsTscAxNotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsTscAxNotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsTscAxNotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

procedure TMsTscAxNotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsTscAxNotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsTscAxNotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsTscAxNotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                           const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

class function CoMsTscAx.Create: IMsTscAx;
begin
  Result := CreateComObject(CLASS_MsTscAx) as IMsTscAx;
end;

class function CoMsTscAx.CreateRemote(const MachineName: string): IMsTscAx;
begin
  Result := CreateRemoteComObject(MachineName, CLASS_MsTscAx) as IMsTscAx;
end;

procedure TMsTscAx.InitServerData;
const
  CServerData: TServerData = (
    ClassID:   '{1FB464C8-09BB-4017-A2F5-EB742F04392F}';
    IntfIID:   '{8C11EFAE-92C3-11D1-BC1E-00C04FA31489}';
    EventIID:  '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    LicenseKey: nil;
    Version: 500);
begin
  ServerData := @CServerData;
end;

procedure TMsTscAx.Connect;
var
  punk: IUnknown;
begin
  if FIntf = nil then
  begin
    punk := GetServer;
    ConnectEvents(punk);
    Fintf:= punk as IMsTscAx;
  end;
end;

procedure TMsTscAx.ConnectTo(svrIntf: IMsTscAx);
begin
  Disconnect;
  FIntf := svrIntf;
  ConnectEvents(FIntf);
end;

procedure TMsTscAx.DisConnect;
begin
  if Fintf <> nil then
  begin
    DisconnectEvents(FIntf);
    FIntf := nil;
  end;
end;

function TMsTscAx.GetDefaultInterface: IMsTscAx;
begin
  if FIntf = nil then
    Connect;
  Assert(FIntf <> nil, 'DefaultInterface is NULL. Component is not connected to Server. You must call "Connect" or "ConnectTo" before this operation');
  Result := FIntf;
end;

constructor TMsTscAx.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
end;

destructor TMsTscAx.Destroy;
begin
  inherited Destroy;
end;

procedure TMsTscAx.InvokeEvent(DispID: TDispID; var Params: TVariantArray);
begin
  case DispID of
    -1: Exit;  // DISPID_UNKNOWN
    1: if Assigned(FOnConnecting) then
         FOnConnecting(Self);
    2: if Assigned(FOnConnected) then
         FOnConnected(Self);
    3: if Assigned(FOnLoginComplete) then
         FOnLoginComplete(Self);
    4: if Assigned(FOnDisconnected) then
         FOnDisconnected(Self, Params[0] {Integer});
    5: if Assigned(FOnEnterFullScreenMode) then
         FOnEnterFullScreenMode(Self);
    6: if Assigned(FOnLeaveFullScreenMode) then
         FOnLeaveFullScreenMode(Self);
    7: if Assigned(FOnChannelReceivedData) then
         FOnChannelReceivedData(Self,
                                Params[0] {const WideString},
                                Params[1] {const WideString});
    8: if Assigned(FOnRequestGoFullScreen) then
         FOnRequestGoFullScreen(Self);
    9: if Assigned(FOnRequestLeaveFullScreen) then
         FOnRequestLeaveFullScreen(Self);
    10: if Assigned(FOnFatalError) then
         FOnFatalError(Self, Params[0] {Integer});
    11: if Assigned(FOnWarning) then
         FOnWarning(Self, Params[0] {Integer});
    12: if Assigned(FOnRemoteDesktopSizeChange) then
         FOnRemoteDesktopSizeChange(Self,
                                    Params[0] {Integer},
                                    Params[1] {Integer});
    13: if Assigned(FOnIdleTimeoutNotification) then
         FOnIdleTimeoutNotification(Self);
    14: if Assigned(FOnRequestContainerMinimize) then
         FOnRequestContainerMinimize(Self);
    15: if Assigned(FOnConfirmClose) then
         FOnConfirmClose(Self);
    16: if Assigned(FOnReceivedTSPublicKey) then
         FOnReceivedTSPublicKey(Self, Params[0] {const WideString});
    17: if Assigned(FOnAutoReconnecting) then
         FOnAutoReconnecting(Self,
                             Params[0] {Integer},
                             Params[1] {Integer});
    18: if Assigned(FOnAuthenticationWarningDisplayed) then
         FOnAuthenticationWarningDisplayed(Self);
    19: if Assigned(FOnAuthenticationWarningDismissed) then
         FOnAuthenticationWarningDismissed(Self);
    20: if Assigned(FOnRemoteProgramResult) then
         FOnRemoteProgramResult(Self,
                                Params[0] {const WideString},
                                Params[1] {RemoteProgramResult},
                                Params[2] {WordBool});
    21: if Assigned(FOnRemoteProgramDisplayed) then
         FOnRemoteProgramDisplayed(Self,
                                   Params[0] {WordBool},
                                   Params[1] {LongWord});
    29: if Assigned(FOnRemoteWindowDisplayed) then
         FOnRemoteWindowDisplayed(Self,
                                  Params[0] {WordBool},
                                  {NOT_OLEAUTO(_RemotableHandle)}OleVariant((TVarData(Params[1]).VPointer)^) {var  NOT_OLEAUTO(_RemotableHandle) OleVariant},
                                  Params[2] {RemoteWindowDisplayedAttribute});
    22: if Assigned(FOnLogonError) then
         FOnLogonError(Self, Params[0] {Integer});
    23: if Assigned(FOnFocusReleased) then
         FOnFocusReleased(Self, Params[0] {SYSINT});
    24: if Assigned(FOnUserNameAcquired) then
         FOnUserNameAcquired(Self, Params[0] {const WideString});
    26: if Assigned(FOnMouseInputModeChanged) then
         FOnMouseInputModeChanged(Self, Params[0] {WordBool});
    28: if Assigned(FOnServiceMessageReceived) then
         FOnServiceMessageReceived(Self, Params[0] {const WideString});
    30: if Assigned(FOnConnectionBarPullDown) then
         FOnConnectionBarPullDown(Self);
    32: if Assigned(FOnNetworkStatusChanged) then
         FOnNetworkStatusChanged(Self,
                                 Params[0] {LongWord},
                                 Params[1] {Integer},
                                 Params[2] {Integer});
    35: if Assigned(FOnDevicesButtonPressed) then
         FOnDevicesButtonPressed(Self);
    33: if Assigned(FOnAutoReconnected) then
         FOnAutoReconnected(Self);
    34: if Assigned(FOnAutoReconnecting2) then
         FOnAutoReconnecting2(Self,
                              Params[0] {Integer},
                              Params[1] {WordBool},
                              Params[2] {Integer},
                              Params[3] {Integer});
  end; {case DispID}
end;

procedure TMsTscAx.Set_Server(const pServer: WideString);
begin
  DefaultInterface.Server := pServer;
end;

function TMsTscAx.Get_Server: WideString;
begin
  Result := DefaultInterface.Server;
end;

procedure TMsTscAx.Set_Domain(const pDomain: WideString);
begin
  DefaultInterface.Domain := pDomain;
end;

function TMsTscAx.Get_Domain: WideString;
begin
  Result := DefaultInterface.Domain;
end;

procedure TMsTscAx.Set_UserName(const pUserName: WideString);
begin
  DefaultInterface.UserName := pUserName;
end;

function TMsTscAx.Get_UserName: WideString;
begin
  Result := DefaultInterface.UserName;
end;

procedure TMsTscAx.Set_DisconnectedText(const pDisconnectedText: WideString);
begin
  DefaultInterface.DisconnectedText := pDisconnectedText;
end;

function TMsTscAx.Get_DisconnectedText: WideString;
begin
  Result := DefaultInterface.DisconnectedText;
end;

procedure TMsTscAx.Set_ConnectingText(const pConnectingText: WideString);
begin
  DefaultInterface.ConnectingText := pConnectingText;
end;

function TMsTscAx.Get_ConnectingText: WideString;
begin
  Result := DefaultInterface.ConnectingText;
end;

function TMsTscAx.Get_Connected: Smallint;
begin
  Result := DefaultInterface.Connected;
end;

procedure TMsTscAx.Set_DesktopWidth(pVal: Integer);
begin
  DefaultInterface.DesktopWidth := pVal;
end;

function TMsTscAx.Get_DesktopWidth: Integer;
begin
  Result := DefaultInterface.DesktopWidth;
end;

procedure TMsTscAx.Set_DesktopHeight(pVal: Integer);
begin
  DefaultInterface.DesktopHeight := pVal;
end;

function TMsTscAx.Get_DesktopHeight: Integer;
begin
  Result := DefaultInterface.DesktopHeight;
end;

procedure TMsTscAx.Set_StartConnected(pfStartConnected: Integer);
begin
  DefaultInterface.StartConnected := pfStartConnected;
end;

function TMsTscAx.Get_StartConnected: Integer;
begin
  Result := DefaultInterface.StartConnected;
end;

function TMsTscAx.Get_HorizontalScrollBarVisible: Integer;
begin
  Result := DefaultInterface.HorizontalScrollBarVisible;
end;

function TMsTscAx.Get_VerticalScrollBarVisible: Integer;
begin
  Result := DefaultInterface.VerticalScrollBarVisible;
end;

procedure TMsTscAx.Set_FullScreenTitle(const Param1: WideString);
begin
  DefaultInterface.FullScreenTitle := Param1;
end;

function TMsTscAx.Get_CipherStrength: Integer;
begin
  Result := DefaultInterface.CipherStrength;
end;

function TMsTscAx.Get_Version: WideString;
begin
  Result := DefaultInterface.Version;
end;

function TMsTscAx.Get_SecuredSettingsEnabled: Integer;
begin
  Result := DefaultInterface.SecuredSettingsEnabled;
end;

function TMsTscAx.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsTscAx.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsTscAx.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

procedure TMsTscAx.Connect1;
begin
  DefaultInterface.Connect;
end;

procedure TMsTscAx.Disconnect1;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsTscAx.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsTscAx.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClientNotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{7CACBD7B-0D99-468F-AC33-22E495C0AFE5}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClientNotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClientNotSafeForScripting.GetControlInterface: IMsRdpClient;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClientNotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClientNotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClientNotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClientNotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClientNotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

procedure TMsRdpClientNotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClientNotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClientNotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClientNotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                               const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClientNotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                   chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClientNotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClientNotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

class function CoMsRdpClient.Create: IMsRdpClient;
begin
  Result := CreateComObject(CLASS_MsRdpClient) as IMsRdpClient;
end;

class function CoMsRdpClient.CreateRemote(const MachineName: string): IMsRdpClient;
begin
  Result := CreateRemoteComObject(MachineName, CLASS_MsRdpClient) as IMsRdpClient;
end;

procedure TMsRdpClient.InitServerData;
const
  CServerData: TServerData = (
    ClassID:   '{791FA017-2DE3-492E-ACC5-53C67A2B94D0}';
    IntfIID:   '{92B4A539-7115-4B7C-A5A9-E5D9EFC2780A}';
    EventIID:  '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    LicenseKey: nil;
    Version: 500);
begin
  ServerData := @CServerData;
end;

procedure TMsRdpClient.Connect;
var
  punk: IUnknown;
begin
  if FIntf = nil then
  begin
    punk := GetServer;
    ConnectEvents(punk);
    Fintf:= punk as IMsRdpClient;
  end;
end;

procedure TMsRdpClient.ConnectTo(svrIntf: IMsRdpClient);
begin
  Disconnect;
  FIntf := svrIntf;
  ConnectEvents(FIntf);
end;

procedure TMsRdpClient.DisConnect;
begin
  if Fintf <> nil then
  begin
    DisconnectEvents(FIntf);
    FIntf := nil;
  end;
end;

function TMsRdpClient.GetDefaultInterface: IMsRdpClient;
begin
  if FIntf = nil then
    Connect;
  Assert(FIntf <> nil, 'DefaultInterface is NULL. Component is not connected to Server. You must call "Connect" or "ConnectTo" before this operation');
  Result := FIntf;
end;

constructor TMsRdpClient.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
end;

destructor TMsRdpClient.Destroy;
begin
  inherited Destroy;
end;

procedure TMsRdpClient.InvokeEvent(DispID: TDispID; var Params: TVariantArray);
begin
  case DispID of
    -1: Exit;  // DISPID_UNKNOWN
    1: if Assigned(FOnConnecting) then
         FOnConnecting(Self);
    2: if Assigned(FOnConnected) then
         FOnConnected(Self);
    3: if Assigned(FOnLoginComplete) then
         FOnLoginComplete(Self);
    4: if Assigned(FOnDisconnected) then
         FOnDisconnected(Self, Params[0] {Integer});
    5: if Assigned(FOnEnterFullScreenMode) then
         FOnEnterFullScreenMode(Self);
    6: if Assigned(FOnLeaveFullScreenMode) then
         FOnLeaveFullScreenMode(Self);
    7: if Assigned(FOnChannelReceivedData) then
         FOnChannelReceivedData(Self,
                                Params[0] {const WideString},
                                Params[1] {const WideString});
    8: if Assigned(FOnRequestGoFullScreen) then
         FOnRequestGoFullScreen(Self);
    9: if Assigned(FOnRequestLeaveFullScreen) then
         FOnRequestLeaveFullScreen(Self);
    10: if Assigned(FOnFatalError) then
         FOnFatalError(Self, Params[0] {Integer});
    11: if Assigned(FOnWarning) then
         FOnWarning(Self, Params[0] {Integer});
    12: if Assigned(FOnRemoteDesktopSizeChange) then
         FOnRemoteDesktopSizeChange(Self,
                                    Params[0] {Integer},
                                    Params[1] {Integer});
    13: if Assigned(FOnIdleTimeoutNotification) then
         FOnIdleTimeoutNotification(Self);
    14: if Assigned(FOnRequestContainerMinimize) then
         FOnRequestContainerMinimize(Self);
    15: if Assigned(FOnConfirmClose) then
         FOnConfirmClose(Self);
    16: if Assigned(FOnReceivedTSPublicKey) then
         FOnReceivedTSPublicKey(Self, Params[0] {const WideString});
    17: if Assigned(FOnAutoReconnecting) then
         FOnAutoReconnecting(Self,
                             Params[0] {Integer},
                             Params[1] {Integer});
    18: if Assigned(FOnAuthenticationWarningDisplayed) then
         FOnAuthenticationWarningDisplayed(Self);
    19: if Assigned(FOnAuthenticationWarningDismissed) then
         FOnAuthenticationWarningDismissed(Self);
    20: if Assigned(FOnRemoteProgramResult) then
         FOnRemoteProgramResult(Self,
                                Params[0] {const WideString},
                                Params[1] {RemoteProgramResult},
                                Params[2] {WordBool});
    21: if Assigned(FOnRemoteProgramDisplayed) then
         FOnRemoteProgramDisplayed(Self,
                                   Params[0] {WordBool},
                                   Params[1] {LongWord});
    29: if Assigned(FOnRemoteWindowDisplayed) then
         FOnRemoteWindowDisplayed(Self,
                                  Params[0] {WordBool},
                                  {NOT_OLEAUTO(_RemotableHandle)}OleVariant((TVarData(Params[1]).VPointer)^) {var  NOT_OLEAUTO(_RemotableHandle) OleVariant},
                                  Params[2] {RemoteWindowDisplayedAttribute});
    22: if Assigned(FOnLogonError) then
         FOnLogonError(Self, Params[0] {Integer});
    23: if Assigned(FOnFocusReleased) then
         FOnFocusReleased(Self, Params[0] {SYSINT});
    24: if Assigned(FOnUserNameAcquired) then
         FOnUserNameAcquired(Self, Params[0] {const WideString});
    26: if Assigned(FOnMouseInputModeChanged) then
         FOnMouseInputModeChanged(Self, Params[0] {WordBool});
    28: if Assigned(FOnServiceMessageReceived) then
         FOnServiceMessageReceived(Self, Params[0] {const WideString});
    30: if Assigned(FOnConnectionBarPullDown) then
         FOnConnectionBarPullDown(Self);
    32: if Assigned(FOnNetworkStatusChanged) then
         FOnNetworkStatusChanged(Self,
                                 Params[0] {LongWord},
                                 Params[1] {Integer},
                                 Params[2] {Integer});
    35: if Assigned(FOnDevicesButtonPressed) then
         FOnDevicesButtonPressed(Self);
    33: if Assigned(FOnAutoReconnected) then
         FOnAutoReconnected(Self);
    34: if Assigned(FOnAutoReconnecting2) then
         FOnAutoReconnecting2(Self,
                              Params[0] {Integer},
                              Params[1] {WordBool},
                              Params[2] {Integer},
                              Params[3] {Integer});
  end; {case DispID}
end;

procedure TMsRdpClient.Set_Server(const pServer: WideString);
begin
  DefaultInterface.Server := pServer;
end;

function TMsRdpClient.Get_Server: WideString;
begin
  Result := DefaultInterface.Server;
end;

procedure TMsRdpClient.Set_Domain(const pDomain: WideString);
begin
  DefaultInterface.Domain := pDomain;
end;

function TMsRdpClient.Get_Domain: WideString;
begin
  Result := DefaultInterface.Domain;
end;

procedure TMsRdpClient.Set_UserName(const pUserName: WideString);
begin
  DefaultInterface.UserName := pUserName;
end;

function TMsRdpClient.Get_UserName: WideString;
begin
  Result := DefaultInterface.UserName;
end;

procedure TMsRdpClient.Set_DisconnectedText(const pDisconnectedText: WideString);
begin
  DefaultInterface.DisconnectedText := pDisconnectedText;
end;

function TMsRdpClient.Get_DisconnectedText: WideString;
begin
  Result := DefaultInterface.DisconnectedText;
end;

procedure TMsRdpClient.Set_ConnectingText(const pConnectingText: WideString);
begin
  DefaultInterface.ConnectingText := pConnectingText;
end;

function TMsRdpClient.Get_ConnectingText: WideString;
begin
  Result := DefaultInterface.ConnectingText;
end;

function TMsRdpClient.Get_Connected: Smallint;
begin
  Result := DefaultInterface.Connected;
end;

procedure TMsRdpClient.Set_DesktopWidth(pVal: Integer);
begin
  DefaultInterface.DesktopWidth := pVal;
end;

function TMsRdpClient.Get_DesktopWidth: Integer;
begin
  Result := DefaultInterface.DesktopWidth;
end;

procedure TMsRdpClient.Set_DesktopHeight(pVal: Integer);
begin
  DefaultInterface.DesktopHeight := pVal;
end;

function TMsRdpClient.Get_DesktopHeight: Integer;
begin
  Result := DefaultInterface.DesktopHeight;
end;

procedure TMsRdpClient.Set_StartConnected(pfStartConnected: Integer);
begin
  DefaultInterface.StartConnected := pfStartConnected;
end;

function TMsRdpClient.Get_StartConnected: Integer;
begin
  Result := DefaultInterface.StartConnected;
end;

function TMsRdpClient.Get_HorizontalScrollBarVisible: Integer;
begin
  Result := DefaultInterface.HorizontalScrollBarVisible;
end;

function TMsRdpClient.Get_VerticalScrollBarVisible: Integer;
begin
  Result := DefaultInterface.VerticalScrollBarVisible;
end;

procedure TMsRdpClient.Set_FullScreenTitle(const Param1: WideString);
begin
  DefaultInterface.FullScreenTitle := Param1;
end;

function TMsRdpClient.Get_CipherStrength: Integer;
begin
  Result := DefaultInterface.CipherStrength;
end;

function TMsRdpClient.Get_Version: WideString;
begin
  Result := DefaultInterface.Version;
end;

function TMsRdpClient.Get_SecuredSettingsEnabled: Integer;
begin
  Result := DefaultInterface.SecuredSettingsEnabled;
end;

function TMsRdpClient.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

procedure TMsRdpClient.Set_ColorDepth(pcolorDepth: Integer);
begin
  DefaultInterface.ColorDepth := pcolorDepth;
end;

function TMsRdpClient.Get_ColorDepth: Integer;
begin
  Result := DefaultInterface.ColorDepth;
end;

function TMsRdpClient.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient.Get_ExtendedDisconnectReason: ExtendedDisconnectReasonCode;
begin
  Result := DefaultInterface.ExtendedDisconnectReason;
end;

procedure TMsRdpClient.Set_FullScreen(pfFullScreen: WordBool);
begin
  DefaultInterface.FullScreen := pfFullScreen;
end;

function TMsRdpClient.Get_FullScreen: WordBool;
begin
  Result := DefaultInterface.FullScreen;
end;

procedure TMsRdpClient.Connect1;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient.Disconnect1;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient2NotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{3523C2FB-4031-44E4-9A3B-F1E94986EE7F}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient2NotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient2;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient2NotSafeForScripting.GetControlInterface: IMsRdpClient2;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient2NotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient2NotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient2NotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient2NotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient2NotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient2NotSafeForScripting.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

procedure TMsRdpClient2NotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient2NotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient2NotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient2NotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                                const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient2NotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                    chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient2NotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient2NotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient2.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{9059F30F-4EB1-4BD2-9FDC-36F43A218F4A}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient2.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient2;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient2.GetControlInterface: IMsRdpClient2;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient2.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient2.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient2.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient2.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient2.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient2.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

procedure TMsRdpClient2.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient2.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient2.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient2.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient2.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient2.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient2.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient2a.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{971127BB-259F-48C2-BD75-5F97A3331551}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient2a.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient2;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient2a.GetControlInterface: IMsRdpClient2;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient2a.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient2a.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient2a.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient2a.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient2a.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient2a.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

procedure TMsRdpClient2a.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient2a.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient2a.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient2a.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient2a.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient2a.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient2a.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient3NotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{ACE575FD-1FCF-4074-9401-EBAB990FA9DE}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient3NotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient3;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient3NotSafeForScripting.GetControlInterface: IMsRdpClient3;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient3NotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient3NotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient3NotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient3NotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient3NotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient3NotSafeForScripting.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient3NotSafeForScripting.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

procedure TMsRdpClient3NotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient3NotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient3NotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient3NotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                                const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient3NotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                    chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient3NotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient3NotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient3.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{7584C670-2274-4EFB-B00B-D6AABA6D3850}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient3.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient3;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient3.GetControlInterface: IMsRdpClient3;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient3.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient3.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient3.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient3.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient3.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient3.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient3.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

procedure TMsRdpClient3.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient3.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient3.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient3.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient3.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient3.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient3.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient3a.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{6A6F4B83-45C5-4CA9-BDD9-0D81C12295E4}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient3a.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient3;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient3a.GetControlInterface: IMsRdpClient3;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient3a.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient3a.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient3a.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient3a.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient3a.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient3a.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient3a.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

procedure TMsRdpClient3a.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient3a.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient3a.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient3a.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient3a.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient3a.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient3a.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient4NotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{6AE29350-321B-42BE-BBE5-12FB5270C0DE}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient4NotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient4;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient4NotSafeForScripting.GetControlInterface: IMsRdpClient4;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient4NotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient4NotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient4NotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient4NotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient4NotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient4NotSafeForScripting.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient4NotSafeForScripting.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient4NotSafeForScripting.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

procedure TMsRdpClient4NotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient4NotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient4NotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient4NotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                                const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient4NotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                    chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient4NotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient4NotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient4.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{4EDCB26C-D24C-4E72-AF07-B576699AC0DE}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient4.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient4;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient4.GetControlInterface: IMsRdpClient4;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient4.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient4.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient4.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient4.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient4.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient4.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient4.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient4.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

procedure TMsRdpClient4.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient4.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient4.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient4.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient4.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient4.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient4.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient4a.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{54CE37E0-9834-41AE-9896-4DAB69DC022B}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient4a.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient4;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient4a.GetControlInterface: IMsRdpClient4;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient4a.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient4a.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient4a.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient4a.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient4a.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient4a.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient4a.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient4a.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

procedure TMsRdpClient4a.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient4a.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient4a.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient4a.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient4a.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient4a.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient4a.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

procedure TMsRdpClient5NotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{4EB2F086-C818-447E-B32C-C51CE2B30D31}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient5NotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient5;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient5NotSafeForScripting.GetControlInterface: IMsRdpClient5;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient5NotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient5NotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient5NotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient5NotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient5NotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient5NotSafeForScripting.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient5NotSafeForScripting.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient5NotSafeForScripting.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient5NotSafeForScripting.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient5NotSafeForScripting.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient5NotSafeForScripting.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient5NotSafeForScripting.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

procedure TMsRdpClient5NotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient5NotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient5NotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient5NotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                                const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient5NotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                    chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient5NotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient5NotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient5NotSafeForScripting.GetErrorDescription(disconnectReason: SYSUINT; 
                                                              ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

procedure TMsRdpClient5.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{4EB89FF4-7F78-4A0F-8B8D-2BF02E94E4B2}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient5.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient5;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient5.GetControlInterface: IMsRdpClient5;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient5.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient5.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient5.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient5.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient5.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient5.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient5.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient5.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient5.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient5.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient5.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient5.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

procedure TMsRdpClient5.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient5.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient5.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient5.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient5.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient5.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient5.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient5.GetErrorDescription(disconnectReason: SYSUINT; 
                                           ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

procedure TMsRdpClient6NotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{D2EA46A7-C2BF-426B-AF24-E19C44456399}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient6NotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient6;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient6NotSafeForScripting.GetControlInterface: IMsRdpClient6;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient6NotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient6NotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient6NotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient6NotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient6NotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient6NotSafeForScripting.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient6NotSafeForScripting.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient6NotSafeForScripting.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient6NotSafeForScripting.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient6NotSafeForScripting.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient6NotSafeForScripting.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient6NotSafeForScripting.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

function TMsRdpClient6NotSafeForScripting.Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
begin
  Result := DefaultInterface.AdvancedSettings7;
end;

function TMsRdpClient6NotSafeForScripting.Get_TransportSettings2: IMsRdpClientTransportSettings2;
begin
  Result := DefaultInterface.TransportSettings2;
end;

procedure TMsRdpClient6NotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient6NotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient6NotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient6NotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                                const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient6NotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                    chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient6NotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient6NotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient6NotSafeForScripting.GetErrorDescription(disconnectReason: SYSUINT; 
                                                              ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

procedure TMsRdpClient6.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{7390F3D8-0439-4C05-91E3-CF5CB290C3D0}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient6.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient6;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient6.GetControlInterface: IMsRdpClient6;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient6.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient6.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient6.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient6.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient6.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient6.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient6.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient6.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient6.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient6.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient6.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient6.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

function TMsRdpClient6.Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
begin
  Result := DefaultInterface.AdvancedSettings7;
end;

function TMsRdpClient6.Get_TransportSettings2: IMsRdpClientTransportSettings2;
begin
  Result := DefaultInterface.TransportSettings2;
end;

procedure TMsRdpClient6.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient6.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient6.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient6.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient6.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient6.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient6.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient6.GetErrorDescription(disconnectReason: SYSUINT; 
                                           ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

procedure TMsRdpClient7NotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{54D38BF7-B1EF-4479-9674-1BD6EA465258}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient7NotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient7;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient7NotSafeForScripting.GetControlInterface: IMsRdpClient7;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient7NotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient7NotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient7NotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient7NotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient7NotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient7NotSafeForScripting.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient7NotSafeForScripting.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient7NotSafeForScripting.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient7NotSafeForScripting.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient7NotSafeForScripting.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient7NotSafeForScripting.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient7NotSafeForScripting.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

function TMsRdpClient7NotSafeForScripting.Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
begin
  Result := DefaultInterface.AdvancedSettings7;
end;

function TMsRdpClient7NotSafeForScripting.Get_TransportSettings2: IMsRdpClientTransportSettings2;
begin
  Result := DefaultInterface.TransportSettings2;
end;

function TMsRdpClient7NotSafeForScripting.Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
begin
  Result := DefaultInterface.AdvancedSettings8;
end;

function TMsRdpClient7NotSafeForScripting.Get_TransportSettings3: IMsRdpClientTransportSettings3;
begin
  Result := DefaultInterface.TransportSettings3;
end;

function TMsRdpClient7NotSafeForScripting.Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
begin
  Result := DefaultInterface.SecuredSettings3;
end;

function TMsRdpClient7NotSafeForScripting.Get_RemoteProgram2: ITSRemoteProgram2;
begin
  Result := DefaultInterface.RemoteProgram2;
end;

procedure TMsRdpClient7NotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient7NotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient7NotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient7NotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                                const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient7NotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                    chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient7NotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient7NotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient7NotSafeForScripting.GetErrorDescription(disconnectReason: SYSUINT; 
                                                              ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

function TMsRdpClient7NotSafeForScripting.GetStatusText(statusCode: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetStatusText(statusCode);
end;

procedure TMsRdpClient7.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{A9D7038D-B5ED-472E-9C47-94BEA90A5910}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient7.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient7;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient7.GetControlInterface: IMsRdpClient7;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient7.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient7.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient7.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient7.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient7.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient7.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient7.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient7.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient7.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient7.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient7.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient7.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

function TMsRdpClient7.Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
begin
  Result := DefaultInterface.AdvancedSettings7;
end;

function TMsRdpClient7.Get_TransportSettings2: IMsRdpClientTransportSettings2;
begin
  Result := DefaultInterface.TransportSettings2;
end;

function TMsRdpClient7.Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
begin
  Result := DefaultInterface.AdvancedSettings8;
end;

function TMsRdpClient7.Get_TransportSettings3: IMsRdpClientTransportSettings3;
begin
  Result := DefaultInterface.TransportSettings3;
end;

function TMsRdpClient7.Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
begin
  Result := DefaultInterface.SecuredSettings3;
end;

function TMsRdpClient7.Get_RemoteProgram2: ITSRemoteProgram2;
begin
  Result := DefaultInterface.RemoteProgram2;
end;

procedure TMsRdpClient7.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient7.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient7.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient7.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient7.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient7.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient7.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient7.GetErrorDescription(disconnectReason: SYSUINT; 
                                           ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

function TMsRdpClient7.GetStatusText(statusCode: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetStatusText(statusCode);
end;

procedure TMsRdpClient8NotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{A3BC03A0-041D-42E3-AD22-882B7865C9C5}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient8NotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient8;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient8NotSafeForScripting.GetControlInterface: IMsRdpClient8;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient8NotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient8NotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient8NotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient8NotSafeForScripting.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient8NotSafeForScripting.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient8NotSafeForScripting.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
begin
  Result := DefaultInterface.AdvancedSettings7;
end;

function TMsRdpClient8NotSafeForScripting.Get_TransportSettings2: IMsRdpClientTransportSettings2;
begin
  Result := DefaultInterface.TransportSettings2;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
begin
  Result := DefaultInterface.AdvancedSettings8;
end;

function TMsRdpClient8NotSafeForScripting.Get_TransportSettings3: IMsRdpClientTransportSettings3;
begin
  Result := DefaultInterface.TransportSettings3;
end;

function TMsRdpClient8NotSafeForScripting.Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
begin
  Result := DefaultInterface.SecuredSettings3;
end;

function TMsRdpClient8NotSafeForScripting.Get_RemoteProgram2: ITSRemoteProgram2;
begin
  Result := DefaultInterface.RemoteProgram2;
end;

function TMsRdpClient8NotSafeForScripting.Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8;
begin
  Result := DefaultInterface.AdvancedSettings9;
end;

procedure TMsRdpClient8NotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient8NotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient8NotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient8NotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                                const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient8NotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                    chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient8NotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient8NotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient8NotSafeForScripting.GetErrorDescription(disconnectReason: SYSUINT; 
                                                              ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

function TMsRdpClient8NotSafeForScripting.GetStatusText(statusCode: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetStatusText(statusCode);
end;

procedure TMsRdpClient8NotSafeForScripting.SendRemoteAction(actionType: RemoteSessionActionType);
begin
  DefaultInterface.SendRemoteAction(actionType);
end;

function TMsRdpClient8NotSafeForScripting.Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus;
begin
  Result := DefaultInterface.Reconnect(ulWidth, ulHeight);
end;

procedure TMsRdpClient8.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{5F681803-2900-4C43-A1CC-CF405404A676}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient8.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient8;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient8.GetControlInterface: IMsRdpClient8;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient8.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient8.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient8.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient8.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient8.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient8.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient8.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient8.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient8.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient8.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient8.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient8.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

function TMsRdpClient8.Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
begin
  Result := DefaultInterface.AdvancedSettings7;
end;

function TMsRdpClient8.Get_TransportSettings2: IMsRdpClientTransportSettings2;
begin
  Result := DefaultInterface.TransportSettings2;
end;

function TMsRdpClient8.Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
begin
  Result := DefaultInterface.AdvancedSettings8;
end;

function TMsRdpClient8.Get_TransportSettings3: IMsRdpClientTransportSettings3;
begin
  Result := DefaultInterface.TransportSettings3;
end;

function TMsRdpClient8.Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
begin
  Result := DefaultInterface.SecuredSettings3;
end;

function TMsRdpClient8.Get_RemoteProgram2: ITSRemoteProgram2;
begin
  Result := DefaultInterface.RemoteProgram2;
end;

function TMsRdpClient8.Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8;
begin
  Result := DefaultInterface.AdvancedSettings9;
end;

procedure TMsRdpClient8.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient8.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient8.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient8.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient8.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient8.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient8.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient8.GetErrorDescription(disconnectReason: SYSUINT; 
                                           ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

function TMsRdpClient8.GetStatusText(statusCode: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetStatusText(statusCode);
end;

procedure TMsRdpClient8.SendRemoteAction(actionType: RemoteSessionActionType);
begin
  DefaultInterface.SendRemoteAction(actionType);
end;

function TMsRdpClient8.Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus;
begin
  Result := DefaultInterface.Reconnect(ulWidth, ulHeight);
end;

procedure TMsRdpClient9NotSafeForScripting.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{8B918B82-7985-4C24-89DF-C33AD2BBFBCD}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient9NotSafeForScripting.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient9;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient9NotSafeForScripting.GetControlInterface: IMsRdpClient9;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient9NotSafeForScripting.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient9NotSafeForScripting.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient9NotSafeForScripting.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient9NotSafeForScripting.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient9NotSafeForScripting.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient9NotSafeForScripting.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
begin
  Result := DefaultInterface.AdvancedSettings7;
end;

function TMsRdpClient9NotSafeForScripting.Get_TransportSettings2: IMsRdpClientTransportSettings2;
begin
  Result := DefaultInterface.TransportSettings2;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
begin
  Result := DefaultInterface.AdvancedSettings8;
end;

function TMsRdpClient9NotSafeForScripting.Get_TransportSettings3: IMsRdpClientTransportSettings3;
begin
  Result := DefaultInterface.TransportSettings3;
end;

function TMsRdpClient9NotSafeForScripting.Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
begin
  Result := DefaultInterface.SecuredSettings3;
end;

function TMsRdpClient9NotSafeForScripting.Get_RemoteProgram2: ITSRemoteProgram2;
begin
  Result := DefaultInterface.RemoteProgram2;
end;

function TMsRdpClient9NotSafeForScripting.Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8;
begin
  Result := DefaultInterface.AdvancedSettings9;
end;

function TMsRdpClient9NotSafeForScripting.Get_TransportSettings4: IMsRdpClientTransportSettings4;
begin
  Result := DefaultInterface.TransportSettings4;
end;

procedure TMsRdpClient9NotSafeForScripting.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient9NotSafeForScripting.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient9NotSafeForScripting.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient9NotSafeForScripting.SendOnVirtualChannel(const chanName: WideString; 
                                                                const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient9NotSafeForScripting.SetVirtualChannelOptions(const chanName: WideString; 
                                                                    chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient9NotSafeForScripting.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient9NotSafeForScripting.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient9NotSafeForScripting.GetErrorDescription(disconnectReason: SYSUINT; 
                                                              ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

function TMsRdpClient9NotSafeForScripting.GetStatusText(statusCode: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetStatusText(statusCode);
end;

procedure TMsRdpClient9NotSafeForScripting.SendRemoteAction(actionType: RemoteSessionActionType);
begin
  DefaultInterface.SendRemoteAction(actionType);
end;

function TMsRdpClient9NotSafeForScripting.Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus;
begin
  Result := DefaultInterface.Reconnect(ulWidth, ulHeight);
end;

procedure TMsRdpClient9NotSafeForScripting.SyncSessionDisplaySettings;
begin
  DefaultInterface.SyncSessionDisplaySettings;
end;

procedure TMsRdpClient9NotSafeForScripting.UpdateSessionDisplaySettings(ulDesktopWidth: LongWord; 
                                                                        ulDesktopHeight: LongWord; 
                                                                        ulPhysicalWidth: LongWord; 
                                                                        ulPhysicalHeight: LongWord; 
                                                                        ulOrientation: LongWord; 
                                                                        ulDesktopScaleFactor: LongWord; 
                                                                        ulDeviceScaleFactor: LongWord);
begin
  DefaultInterface.UpdateSessionDisplaySettings(ulDesktopWidth, ulDesktopHeight, ulPhysicalWidth, 
                                                ulPhysicalHeight, ulOrientation, 
                                                ulDesktopScaleFactor, ulDeviceScaleFactor);
end;

procedure TMsRdpClient9NotSafeForScripting.attachEvent(const eventName: WideString; 
                                                       const callback: IDispatch);
begin
  DefaultInterface.attachEvent(eventName, callback);
end;

procedure TMsRdpClient9NotSafeForScripting.detachEvent(const eventName: WideString; 
                                                       const callback: IDispatch);
begin
  DefaultInterface.detachEvent(eventName, callback);
end;

procedure TMsRdpClient9.InitControlData;
const
  CEventDispIDs: array [0..31] of DWORD = (
    $00000001, $00000002, $00000003, $00000004, $00000005, $00000006,
    $00000007, $00000008, $00000009, $0000000A, $0000000B, $0000000C,
    $0000000D, $0000000E, $0000000F, $00000010, $00000011, $00000012,
    $00000013, $00000014, $00000015, $0000001D, $00000016, $00000017,
    $00000018, $0000001A, $0000001C, $0000001E, $00000020, $00000023,
    $00000021, $00000022);
  CControlData: TControlData2 = (
    ClassID:      '{301B94BA-5D25-4A12-BFFE-3B6E7A616585}';
    EventIID:     '{336D5562-EFA8-482E-8CB3-C5C0FC7A7DB6}';
    EventCount:   32;
    EventDispIDs: @CEventDispIDs;
    LicenseKey:   nil (*HR:$80040111*);
    Flags:        $00000000;
    Version:      500);
begin
  ControlData := @CControlData;
  TControlData2(CControlData).FirstEventOfs := UIntPtr(@@FOnConnecting) - UIntPtr(Self);
end;

procedure TMsRdpClient9.CreateControl;

  procedure DoCreate;
  begin
    FIntf := IUnknown(OleObject) as IMsRdpClient9;
  end;

begin
  if FIntf = nil then DoCreate;
end;

function TMsRdpClient9.GetControlInterface: IMsRdpClient9;
begin
  CreateControl;
  Result := FIntf;
end;

function TMsRdpClient9.Get_SecuredSettings: IMsTscSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings;
end;

function TMsRdpClient9.Get_AdvancedSettings: IMsTscAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings;
end;

function TMsRdpClient9.Get_Debugger: IMsTscDebug;
begin
  Result := DefaultInterface.Debugger;
end;

function TMsRdpClient9.Get_AdvancedSettings2: IMsRdpClientAdvancedSettings;
begin
  Result := DefaultInterface.AdvancedSettings2;
end;

function TMsRdpClient9.Get_SecuredSettings2: IMsRdpClientSecuredSettings;
begin
  Result := DefaultInterface.SecuredSettings2;
end;

function TMsRdpClient9.Get_AdvancedSettings3: IMsRdpClientAdvancedSettings2;
begin
  Result := DefaultInterface.AdvancedSettings3;
end;

function TMsRdpClient9.Get_AdvancedSettings4: IMsRdpClientAdvancedSettings3;
begin
  Result := DefaultInterface.AdvancedSettings4;
end;

function TMsRdpClient9.Get_AdvancedSettings5: IMsRdpClientAdvancedSettings4;
begin
  Result := DefaultInterface.AdvancedSettings5;
end;

function TMsRdpClient9.Get_TransportSettings: IMsRdpClientTransportSettings;
begin
  Result := DefaultInterface.TransportSettings;
end;

function TMsRdpClient9.Get_AdvancedSettings6: IMsRdpClientAdvancedSettings5;
begin
  Result := DefaultInterface.AdvancedSettings6;
end;

function TMsRdpClient9.Get_RemoteProgram: ITSRemoteProgram;
begin
  Result := DefaultInterface.RemoteProgram;
end;

function TMsRdpClient9.Get_MsRdpClientShell: IMsRdpClientShell;
begin
  Result := DefaultInterface.MsRdpClientShell;
end;

function TMsRdpClient9.Get_AdvancedSettings7: IMsRdpClientAdvancedSettings6;
begin
  Result := DefaultInterface.AdvancedSettings7;
end;

function TMsRdpClient9.Get_TransportSettings2: IMsRdpClientTransportSettings2;
begin
  Result := DefaultInterface.TransportSettings2;
end;

function TMsRdpClient9.Get_AdvancedSettings8: IMsRdpClientAdvancedSettings7;
begin
  Result := DefaultInterface.AdvancedSettings8;
end;

function TMsRdpClient9.Get_TransportSettings3: IMsRdpClientTransportSettings3;
begin
  Result := DefaultInterface.TransportSettings3;
end;

function TMsRdpClient9.Get_SecuredSettings3: IMsRdpClientSecuredSettings2;
begin
  Result := DefaultInterface.SecuredSettings3;
end;

function TMsRdpClient9.Get_RemoteProgram2: ITSRemoteProgram2;
begin
  Result := DefaultInterface.RemoteProgram2;
end;

function TMsRdpClient9.Get_AdvancedSettings9: IMsRdpClientAdvancedSettings8;
begin
  Result := DefaultInterface.AdvancedSettings9;
end;

function TMsRdpClient9.Get_TransportSettings4: IMsRdpClientTransportSettings4;
begin
  Result := DefaultInterface.TransportSettings4;
end;

procedure TMsRdpClient9.Connect;
begin
  DefaultInterface.Connect;
end;

procedure TMsRdpClient9.Disconnect;
begin
  DefaultInterface.Disconnect;
end;

procedure TMsRdpClient9.CreateVirtualChannels(const newVal: WideString);
begin
  DefaultInterface.CreateVirtualChannels(newVal);
end;

procedure TMsRdpClient9.SendOnVirtualChannel(const chanName: WideString; const ChanData: WideString);
begin
  DefaultInterface.SendOnVirtualChannel(chanName, ChanData);
end;

procedure TMsRdpClient9.SetVirtualChannelOptions(const chanName: WideString; chanOptions: Integer);
begin
  DefaultInterface.SetVirtualChannelOptions(chanName, chanOptions);
end;

function TMsRdpClient9.GetVirtualChannelOptions(const chanName: WideString): Integer;
begin
  Result := DefaultInterface.GetVirtualChannelOptions(chanName);
end;

function TMsRdpClient9.RequestClose: ControlCloseStatus;
begin
  Result := DefaultInterface.RequestClose;
end;

function TMsRdpClient9.GetErrorDescription(disconnectReason: SYSUINT; 
                                           ExtendedDisconnectReason: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetErrorDescription(disconnectReason, ExtendedDisconnectReason);
end;

function TMsRdpClient9.GetStatusText(statusCode: SYSUINT): WideString;
begin
  Result := DefaultInterface.GetStatusText(statusCode);
end;

procedure TMsRdpClient9.SendRemoteAction(actionType: RemoteSessionActionType);
begin
  DefaultInterface.SendRemoteAction(actionType);
end;

function TMsRdpClient9.Reconnect(ulWidth: LongWord; ulHeight: LongWord): ControlReconnectStatus;
begin
  Result := DefaultInterface.Reconnect(ulWidth, ulHeight);
end;

procedure TMsRdpClient9.SyncSessionDisplaySettings;
begin
  DefaultInterface.SyncSessionDisplaySettings;
end;

procedure TMsRdpClient9.UpdateSessionDisplaySettings(ulDesktopWidth: LongWord; 
                                                     ulDesktopHeight: LongWord; 
                                                     ulPhysicalWidth: LongWord; 
                                                     ulPhysicalHeight: LongWord; 
                                                     ulOrientation: LongWord; 
                                                     ulDesktopScaleFactor: LongWord; 
                                                     ulDeviceScaleFactor: LongWord);
begin
  DefaultInterface.UpdateSessionDisplaySettings(ulDesktopWidth, ulDesktopHeight, ulPhysicalWidth, 
                                                ulPhysicalHeight, ulOrientation, 
                                                ulDesktopScaleFactor, ulDeviceScaleFactor);
end;

procedure TMsRdpClient9.attachEvent(const eventName: WideString; const callback: IDispatch);
begin
  DefaultInterface.attachEvent(eventName, callback);
end;

procedure TMsRdpClient9.detachEvent(const eventName: WideString; const callback: IDispatch);
begin
  DefaultInterface.detachEvent(eventName, callback);
end;

class function CoRemoteDesktopClient.Create: IRemoteDesktopClient;
begin
  Result := CreateComObject(CLASS_RemoteDesktopClient) as IRemoteDesktopClient;
end;

class function CoRemoteDesktopClient.CreateRemote(const MachineName: string): IRemoteDesktopClient;
begin
  Result := CreateRemoteComObject(MachineName, CLASS_RemoteDesktopClient) as IRemoteDesktopClient;
end;

procedure TRemoteDesktopClient.InitServerData;
const
  CServerData: TServerData = (
    ClassID:   '{EAB16C5D-EED1-4E95-868B-0FBA1B42C092}';
    IntfIID:   '{57D25668-625A-4905-BE4E-304CAA13F89C}';
    EventIID:  '{079863B7-6D47-4105-8BFE-0CDCB360E67D}';
    LicenseKey: nil;
    Version: 500);
begin
  ServerData := @CServerData;
end;

procedure TRemoteDesktopClient.Connect;
var
  punk: IUnknown;
begin
  if FIntf = nil then
  begin
    punk := GetServer;
    ConnectEvents(punk);
    Fintf:= punk as IRemoteDesktopClient;
  end;
end;

procedure TRemoteDesktopClient.ConnectTo(svrIntf: IRemoteDesktopClient);
begin
  Disconnect;
  FIntf := svrIntf;
  ConnectEvents(FIntf);
end;

procedure TRemoteDesktopClient.DisConnect;
begin
  if Fintf <> nil then
  begin
    DisconnectEvents(FIntf);
    FIntf := nil;
  end;
end;

function TRemoteDesktopClient.GetDefaultInterface: IRemoteDesktopClient;
begin
  if FIntf = nil then
    Connect;
  Assert(FIntf <> nil, 'DefaultInterface is NULL. Component is not connected to Server. You must call "Connect" or "ConnectTo" before this operation');
  Result := FIntf;
end;

constructor TRemoteDesktopClient.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
end;

destructor TRemoteDesktopClient.Destroy;
begin
  inherited Destroy;
end;

procedure TRemoteDesktopClient.InvokeEvent(DispID: TDispID; var Params: TVariantArray);
begin
  case DispID of
    -1: Exit;  // DISPID_UNKNOWN
    750: if Assigned(FOnConnecting) then
         FOnConnecting(Self);
    751: if Assigned(FOnConnected) then
         FOnConnected(Self);
    752: if Assigned(FOnLoginCompleted) then
         FOnLoginCompleted(Self);
    753: if Assigned(FOnDisconnected) then
         FOnDisconnected(Self,
                         Params[0] {Integer},
                         Params[1] {Integer},
                         Params[2] {const WideString});
    754: if Assigned(FOnStatusChanged) then
         FOnStatusChanged(Self,
                          Params[0] {Integer},
                          Params[1] {const WideString});
    755: if Assigned(FOnAutoReconnecting) then
         FOnAutoReconnecting(Self,
                             Params[0] {Integer},
                             Params[1] {Integer},
                             Params[2] {const WideString},
                             Params[3] {WordBool},
                             Params[4] {Integer},
                             Params[5] {Integer});
    756: if Assigned(FOnAutoReconnected) then
         FOnAutoReconnected(Self);
    757: if Assigned(FOnDialogDisplaying) then
         FOnDialogDisplaying(Self);
    758: if Assigned(FOnDialogDismissed) then
         FOnDialogDismissed(Self);
    759: if Assigned(FOnNetworkStatusChanged) then
         FOnNetworkStatusChanged(Self,
                                 Params[0] {LongWord},
                                 Params[1] {Integer},
                                 Params[2] {Integer});
    760: if Assigned(FOnAdminMessageReceived) then
         FOnAdminMessageReceived(Self, Params[0] {const WideString});
    761: if Assigned(FOnKeyCombinationPressed) then
         FOnKeyCombinationPressed(Self, Params[0] {Integer});
    762: if Assigned(FOnRemoteDesktopSizeChanged) then
         FOnRemoteDesktopSizeChanged(Self,
                                     Params[0] {Integer},
                                     Params[1] {Integer});
    800: if Assigned(FOnTouchPointerCursorMoved) then
         FOnTouchPointerCursorMoved(Self,
                                    Params[0] {Integer},
                                    Params[1] {Integer});
  end; {case DispID}
end;

function TRemoteDesktopClient.Get_Settings: IRemoteDesktopClientSettings;
begin
  Result := DefaultInterface.Settings;
end;

function TRemoteDesktopClient.Get_Actions: IRemoteDesktopClientActions;
begin
  Result := DefaultInterface.Actions;
end;

function TRemoteDesktopClient.Get_TouchPointer: IRemoteDesktopClientTouchPointer;
begin
  Result := DefaultInterface.TouchPointer;
end;

procedure TRemoteDesktopClient.Connect1;
begin
  DefaultInterface.Connect;
end;

procedure TRemoteDesktopClient.Disconnect1;
begin
  DefaultInterface.Disconnect;
end;

procedure TRemoteDesktopClient.Reconnect(width: LongWord; height: LongWord);
begin
  DefaultInterface.Reconnect(width, height);
end;

procedure TRemoteDesktopClient.DeleteSavedCredentials(const serverName: WideString);
begin
  DefaultInterface.DeleteSavedCredentials(serverName);
end;

procedure TRemoteDesktopClient.UpdateSessionDisplaySettings(width: LongWord; height: LongWord);
begin
  DefaultInterface.UpdateSessionDisplaySettings(width, height);
end;

procedure TRemoteDesktopClient.attachEvent(const eventName: WideString; const callback: IDispatch);
begin
  DefaultInterface.attachEvent(eventName, callback);
end;

procedure TRemoteDesktopClient.detachEvent(const eventName: WideString; const callback: IDispatch);
begin
  DefaultInterface.detachEvent(eventName, callback);
end;

procedure Register;
begin
  RegisterComponents(dtlOcxPage, [TMsTscAxNotSafeForScripting, TMsRdpClientNotSafeForScripting, TMsRdpClient2NotSafeForScripting, TMsRdpClient2, 
    TMsRdpClient2a, TMsRdpClient3NotSafeForScripting, TMsRdpClient3, TMsRdpClient3a, TMsRdpClient4NotSafeForScripting, 
    TMsRdpClient4, TMsRdpClient4a, TMsRdpClient5NotSafeForScripting, TMsRdpClient5, TMsRdpClient6NotSafeForScripting, 
    TMsRdpClient6, TMsRdpClient7NotSafeForScripting, TMsRdpClient7, TMsRdpClient8NotSafeForScripting, TMsRdpClient8, 
    TMsRdpClient9NotSafeForScripting, TMsRdpClient9]);
  RegisterComponents(dtlServerPage, [TMsTscAx, TMsRdpClient, TRemoteDesktopClient]);
end;

end.
