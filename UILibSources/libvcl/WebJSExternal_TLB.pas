//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

unit WebJSExternal_TLB;

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
// File generated on 2014/8/11 14:28:33 from Type Library described below.

// ************************************************************************  //
// LIBID: {F3715ED5-C089-4170-94F2-539A125219D7}
// LCID: 0
// Helpfile: 
// HelpString: 
// DepndLst: 
//   (1) v2.0 stdole, (C:\Windows\SysWOW64\stdole2.tlb)
// SYS_KIND: SYS_WIN32
// ************************************************************************ //
{$TYPEDADDRESS OFF} // Unit must be compiled without type-checked pointers. 
{$WARN SYMBOL_PLATFORM OFF}
{$WRITEABLECONST ON}
{$VARPROPSETTER ON}
{$ALIGN 4}

interface

uses Windows, Classes, Variants, Graphics, ActiveX;
  

// *********************************************************************//
// GUIDS declared in the TypeLibrary. Following prefixes are used:        
//   Type Libraries     : LIBID_xxxx                                      
//   CoClasses          : CLASS_xxxx                                      
//   DISPInterfaces     : DIID_xxxx                                       
//   Non-DISP interfaces: IID_xxxx                                        
// *********************************************************************//
const
  // TypeLibrary Major and minor versions
  WebJSExternalMajorVersion = 1;
  WebJSExternalMinorVersion = 0;

  LIBID_WebJSExternal: TGUID = '{F3715ED5-C089-4170-94F2-539A125219D7}';

  IID_IWebJSExternal: TGUID = '{26CE805E-6FF3-420F-B7A6-E8326E4B5CA6}';
type

// *********************************************************************//
// Forward declaration of types defined in TypeLibrary                    
// *********************************************************************//
  IWebJSExternal = interface;
  IWebJSExternalDisp = dispinterface;

// *********************************************************************//
// Interface: IWebJSExternal
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {26CE805E-6FF3-420F-B7A6-E8326E4B5CA6}
// *********************************************************************//
  IWebJSExternal = interface(IDispatch)
    ['{26CE805E-6FF3-420F-B7A6-E8326E4B5CA6}']
    function JSExternal(const func: WideString; const args: WideString): WideString; safecall;
  end;

// *********************************************************************//
// DispIntf:  IWebJSExternalDisp
// Flags:     (4416) Dual OleAutomation Dispatchable
// GUID:      {26CE805E-6FF3-420F-B7A6-E8326E4B5CA6}
// *********************************************************************//
  IWebJSExternalDisp = dispinterface
    ['{26CE805E-6FF3-420F-B7A6-E8326E4B5CA6}']
    function JSExternal(const func: WideString; const args: WideString): WideString; dispid 201;
  end;

implementation


end.
