/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi;

import com.fasterxml.jackson.databind.ObjectMapper;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;
import org.openapis.openapi.utils.HTTPClient;
import org.openapis.openapi.utils.HTTPRequest;
import org.openapis.openapi.utils.JSON;

public class Users {
	
	private HTTPClient _defaultClient;
	private HTTPClient _securityClient;
	private String _serverUrl;
	private String _language;
	private String _sdkVersion;
	private String _genVersion;

	public Users(HTTPClient defaultClient, HTTPClient securityClient, String serverUrl, String language, String sdkVersion, String genVersion) {
		this._defaultClient = defaultClient;
		this._securityClient = securityClient;
		this._serverUrl = serverUrl;
		this._language = language;
		this._sdkVersion = sdkVersion;
		this._genVersion = genVersion;
	}

    /**
     * List users
     * List users
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public org.openapis.openapi.models.operations.ListUsersResponse listUsers() throws Exception {
        String baseUrl = this._serverUrl;
        String url = org.openapis.openapi.utils.Utils.generateURL(baseUrl, "/api/auth/users");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);

        req.addHeader("Accept", "application/json");
        req.addHeader("user-agent", String.format("speakeasy-sdk/%s %s %s", this._language, this._sdkVersion, this._genVersion));
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        org.openapis.openapi.models.operations.ListUsersResponse res = new org.openapis.openapi.models.operations.ListUsersResponse(contentType, httpRes.statusCode()) {{
            listUsersResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (org.openapis.openapi.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                org.openapis.openapi.models.shared.ListUsersResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), org.openapis.openapi.models.shared.ListUsersResponse.class);
                res.listUsersResponse = out;
            }
        }

        return res;
    }

    /**
     * Read user
     * Read user
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public org.openapis.openapi.models.operations.ReadUserResponse readUser(org.openapis.openapi.models.operations.ReadUserRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = org.openapis.openapi.utils.Utils.generateURL(org.openapis.openapi.models.operations.ReadUserRequest.class, baseUrl, "/api/auth/users/{userId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);

        req.addHeader("Accept", "application/json");
        req.addHeader("user-agent", String.format("speakeasy-sdk/%s %s %s", this._language, this._sdkVersion, this._genVersion));
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        org.openapis.openapi.models.operations.ReadUserResponse res = new org.openapis.openapi.models.operations.ReadUserResponse(contentType, httpRes.statusCode()) {{
            readUserResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (org.openapis.openapi.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                org.openapis.openapi.models.shared.ReadUserResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), org.openapis.openapi.models.shared.ReadUserResponse.class);
                res.readUserResponse = out;
            }
        }

        return res;
    }
}