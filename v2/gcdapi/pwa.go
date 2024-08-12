// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains PWA functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// The following types are the replica of https://crsrc.org/c/chrome/browser/web_applications/proto/web_app_os_integration_state.proto;drc=9910d3be894c8f142c977ba1023f30a656bc13fc;l=67
type PWAFileHandlerAccept struct {
	MediaType      string   `json:"mediaType"`      // New name of the mimetype according to https://www.iana.org/assignments/media-types/media-types.xhtml
	FileExtensions []string `json:"fileExtensions"` //
}

// No Description.
type PWAFileHandler struct {
	Action      string                  `json:"action"`      //
	Accepts     []*PWAFileHandlerAccept `json:"accepts"`     //
	DisplayName string                  `json:"displayName"` //
}

type PWA struct {
	target gcdmessage.ChromeTargeter
}

func NewPWA(target gcdmessage.ChromeTargeter) *PWA {
	c := &PWA{target: target}
	return c
}

type PWAGetOsAppStateParams struct {
	// The id from the webapp's manifest file, commonly it's the url of the site installing the webapp. See https://web.dev/learn/pwa/web-app-manifest.
	ManifestId string `json:"manifestId"`
}

// GetOsAppStateWithParams - Returns the following OS state for the given manifest id.
// Returns -  badgeCount -  fileHandlers -
func (c *PWA) GetOsAppStateWithParams(ctx context.Context, v *PWAGetOsAppStateParams) (int, []*PWAFileHandler, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "PWA.getOsAppState", Params: v})
	if err != nil {
		return 0, nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			BadgeCount   int
			FileHandlers []*PWAFileHandler
		}
	}

	if resp == nil {
		return 0, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return 0, nil, err
	}

	if chromeData.Error != nil {
		return 0, nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.BadgeCount, chromeData.Result.FileHandlers, nil
}

// GetOsAppState - Returns the following OS state for the given manifest id.
// manifestId - The id from the webapp's manifest file, commonly it's the url of the site installing the webapp. See https://web.dev/learn/pwa/web-app-manifest.
// Returns -  badgeCount -  fileHandlers -
func (c *PWA) GetOsAppState(ctx context.Context, manifestId string) (int, []*PWAFileHandler, error) {
	var v PWAGetOsAppStateParams
	v.ManifestId = manifestId
	return c.GetOsAppStateWithParams(ctx, &v)
}

type PWAInstallParams struct {
	//
	ManifestId string `json:"manifestId"`
	// The location of the app or bundle overriding the one derived from the manifestId.
	InstallUrlOrBundleUrl string `json:"installUrlOrBundleUrl,omitempty"`
}

// InstallWithParams - Installs the given manifest identity, optionally using the given install_url or IWA bundle location.  TODO(crbug.com/337872319) Support IWA to meet the following specific requirement. IWA-specific install description: If the manifest_id is isolated-app://, install_url_or_bundle_url is required, and can be either an http(s) URL or file:// URL pointing to a signed web bundle (.swbn). The .swbn file's signing key must correspond to manifest_id. If Chrome is not in IWA dev mode, the installation will fail, regardless of the state of the allowlist.
func (c *PWA) InstallWithParams(ctx context.Context, v *PWAInstallParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "PWA.install", Params: v})
}

// Install - Installs the given manifest identity, optionally using the given install_url or IWA bundle location.  TODO(crbug.com/337872319) Support IWA to meet the following specific requirement. IWA-specific install description: If the manifest_id is isolated-app://, install_url_or_bundle_url is required, and can be either an http(s) URL or file:// URL pointing to a signed web bundle (.swbn). The .swbn file's signing key must correspond to manifest_id. If Chrome is not in IWA dev mode, the installation will fail, regardless of the state of the allowlist.
// manifestId -
// installUrlOrBundleUrl - The location of the app or bundle overriding the one derived from the manifestId.
func (c *PWA) Install(ctx context.Context, manifestId string, installUrlOrBundleUrl string) (*gcdmessage.ChromeResponse, error) {
	var v PWAInstallParams
	v.ManifestId = manifestId
	v.InstallUrlOrBundleUrl = installUrlOrBundleUrl
	return c.InstallWithParams(ctx, &v)
}

type PWAUninstallParams struct {
	//
	ManifestId string `json:"manifestId"`
}

// UninstallWithParams - Uninstalls the given manifest_id and closes any opened app windows.
func (c *PWA) UninstallWithParams(ctx context.Context, v *PWAUninstallParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "PWA.uninstall", Params: v})
}

// Uninstall - Uninstalls the given manifest_id and closes any opened app windows.
// manifestId -
func (c *PWA) Uninstall(ctx context.Context, manifestId string) (*gcdmessage.ChromeResponse, error) {
	var v PWAUninstallParams
	v.ManifestId = manifestId
	return c.UninstallWithParams(ctx, &v)
}

type PWALaunchParams struct {
	//
	ManifestId string `json:"manifestId"`
	//
	Url string `json:"url,omitempty"`
}

// LaunchWithParams - Launches the installed web app, or an url in the same web app instead of the default start url if it is provided. Returns a page Target.TargetID which can be used to attach to via Target.attachToTarget or similar APIs.
// Returns -  targetId - ID of the tab target created as a result.
func (c *PWA) LaunchWithParams(ctx context.Context, v *PWALaunchParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "PWA.launch", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			TargetId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	if chromeData.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.TargetId, nil
}

// Launch - Launches the installed web app, or an url in the same web app instead of the default start url if it is provided. Returns a page Target.TargetID which can be used to attach to via Target.attachToTarget or similar APIs.
// manifestId -
// url -
// Returns -  targetId - ID of the tab target created as a result.
func (c *PWA) Launch(ctx context.Context, manifestId string, url string) (string, error) {
	var v PWALaunchParams
	v.ManifestId = manifestId
	v.Url = url
	return c.LaunchWithParams(ctx, &v)
}

type PWALaunchFilesInAppParams struct {
	//
	ManifestId string `json:"manifestId"`
	//
	Files []string `json:"files"`
}

// LaunchFilesInAppWithParams - Opens one or more local files from an installed web app identified by its manifestId. The web app needs to have file handlers registered to process the files. The API returns one or more page Target.TargetIDs which can be used to attach to via Target.attachToTarget or similar APIs. If some files in the parameters cannot be handled by the web app, they will be ignored. If none of the files can be handled, this API returns an error. If no files are provided as the parameter, this API also returns an error.  According to the definition of the file handlers in the manifest file, one Target.TargetID may represent a page handling one or more files. The order of the returned Target.TargetIDs is not guaranteed.  TODO(crbug.com/339454034): Check the existences of the input files.
// Returns -  targetIds - IDs of the tab targets created as the result.
func (c *PWA) LaunchFilesInAppWithParams(ctx context.Context, v *PWALaunchFilesInAppParams) ([]string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "PWA.launchFilesInApp", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		gcdmessage.ChromeErrorResponse
		Result struct {
			TargetIds []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	if err := jsonUnmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	if chromeData.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: &chromeData.ChromeErrorResponse}
	}

	return chromeData.Result.TargetIds, nil
}

// LaunchFilesInApp - Opens one or more local files from an installed web app identified by its manifestId. The web app needs to have file handlers registered to process the files. The API returns one or more page Target.TargetIDs which can be used to attach to via Target.attachToTarget or similar APIs. If some files in the parameters cannot be handled by the web app, they will be ignored. If none of the files can be handled, this API returns an error. If no files are provided as the parameter, this API also returns an error.  According to the definition of the file handlers in the manifest file, one Target.TargetID may represent a page handling one or more files. The order of the returned Target.TargetIDs is not guaranteed.  TODO(crbug.com/339454034): Check the existences of the input files.
// manifestId -
// files -
// Returns -  targetIds - IDs of the tab targets created as the result.
func (c *PWA) LaunchFilesInApp(ctx context.Context, manifestId string, files []string) ([]string, error) {
	var v PWALaunchFilesInAppParams
	v.ManifestId = manifestId
	v.Files = files
	return c.LaunchFilesInAppWithParams(ctx, &v)
}

type PWAOpenCurrentPageInAppParams struct {
	//
	ManifestId string `json:"manifestId"`
}

// OpenCurrentPageInAppWithParams - Opens the current page in its web app identified by the manifest id, needs to be called on a page target. This function returns immediately without waiting for the app to finish loading.
func (c *PWA) OpenCurrentPageInAppWithParams(ctx context.Context, v *PWAOpenCurrentPageInAppParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "PWA.openCurrentPageInApp", Params: v})
}

// OpenCurrentPageInApp - Opens the current page in its web app identified by the manifest id, needs to be called on a page target. This function returns immediately without waiting for the app to finish loading.
// manifestId -
func (c *PWA) OpenCurrentPageInApp(ctx context.Context, manifestId string) (*gcdmessage.ChromeResponse, error) {
	var v PWAOpenCurrentPageInAppParams
	v.ManifestId = manifestId
	return c.OpenCurrentPageInAppWithParams(ctx, &v)
}

type PWAChangeAppUserSettingsParams struct {
	//
	ManifestId string `json:"manifestId"`
	// If user allows the links clicked on by the user in the app's scope, or extended scope if the manifest has scope extensions and the flags `DesktopPWAsLinkCapturingWithScopeExtensions` and `WebAppEnableScopeExtensions` are enabled.  Note, the API does not support resetting the linkCapturing to the initial value, uninstalling and installing the web app again will reset it.  TODO(crbug.com/339453269): Setting this value on ChromeOS is not supported yet.
	LinkCapturing bool `json:"linkCapturing,omitempty"`
	//  enum values: standalone, browser
	DisplayMode string `json:"displayMode,omitempty"`
}

// ChangeAppUserSettingsWithParams - Changes user settings of the web app identified by its manifestId. If the app was not installed, this command returns an error. Unset parameters will be ignored; unrecognized values will cause an error.  Unlike the ones defined in the manifest files of the web apps, these settings are provided by the browser and controlled by the users, they impact the way the browser handling the web apps.  See the comment of each parameter.
func (c *PWA) ChangeAppUserSettingsWithParams(ctx context.Context, v *PWAChangeAppUserSettingsParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "PWA.changeAppUserSettings", Params: v})
}

// ChangeAppUserSettings - Changes user settings of the web app identified by its manifestId. If the app was not installed, this command returns an error. Unset parameters will be ignored; unrecognized values will cause an error.  Unlike the ones defined in the manifest files of the web apps, these settings are provided by the browser and controlled by the users, they impact the way the browser handling the web apps.  See the comment of each parameter.
// manifestId -
// linkCapturing - If user allows the links clicked on by the user in the app's scope, or extended scope if the manifest has scope extensions and the flags `DesktopPWAsLinkCapturingWithScopeExtensions` and `WebAppEnableScopeExtensions` are enabled.  Note, the API does not support resetting the linkCapturing to the initial value, uninstalling and installing the web app again will reset it.  TODO(crbug.com/339453269): Setting this value on ChromeOS is not supported yet.
// displayMode -  enum values: standalone, browser
func (c *PWA) ChangeAppUserSettings(ctx context.Context, manifestId string, linkCapturing bool, displayMode string) (*gcdmessage.ChromeResponse, error) {
	var v PWAChangeAppUserSettingsParams
	v.ManifestId = manifestId
	v.LinkCapturing = linkCapturing
	v.DisplayMode = displayMode
	return c.ChangeAppUserSettingsWithParams(ctx, &v)
}
