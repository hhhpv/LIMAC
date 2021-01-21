/*
 * Copyright 2010-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

var apigClientFactory = {};
apigClientFactory.newClient = function (config) {
    var apigClient = {};
    if (config === undefined) {
        config = {
            accessKey: '',
            secretKey: '',
            sessionToken: '',
            region: '',
            apiKey: undefined,
            defaultContentType: 'application/json',
            defaultAcceptType: 'application/json'
        };
    }
    if (config.accessKey === undefined) {
        config.accessKey = '';
    }
    if (config.secretKey === undefined) {
        config.secretKey = '';
    }
    if (config.apiKey === undefined) {
        config.apiKey = '';
    }
    if (config.sessionToken === undefined) {
        config.sessionToken = '';
    }
    if (config.region === undefined) {
        config.region = 'us-east-1';
    }
    //If defaultContentType is not defined then default to application/json
    if (config.defaultContentType === undefined) {
        config.defaultContentType = 'application/json';
    }
    //If defaultAcceptType is not defined then default to application/json
    if (config.defaultAcceptType === undefined) {
        config.defaultAcceptType = 'application/json';
    }


    // extract endpoint and path from url
    var invokeUrl = 'https://your_endpoint_url.execute-api.us-west-1.amazonaws.com/default';
    var endpoint = /(^https?:\/\/[^\/]+)/g.exec(invokeUrl)[1];
    var pathComponent = invokeUrl.substring(endpoint.length);

    var sigV4ClientConfig = {
        accessKey: config.accessKey,
        secretKey: config.secretKey,
        sessionToken: config.sessionToken,
        serviceName: 'execute-api',
        region: config.region,
        endpoint: endpoint,
        defaultContentType: config.defaultContentType,
        defaultAcceptType: config.defaultAcceptType
    };

    var authType = 'NONE';
    if (sigV4ClientConfig.accessKey !== undefined && sigV4ClientConfig.accessKey !== '' && sigV4ClientConfig.secretKey !== undefined && sigV4ClientConfig.secretKey !== '') {
        authType = 'AWS_IAM';
    }

    var simpleHttpClientConfig = {
        endpoint: endpoint,
        defaultContentType: config.defaultContentType,
        defaultAcceptType: config.defaultAcceptType
    };

    var apiGatewayClient = apiGateway.core.apiGatewayClientFactory.newClient(simpleHttpClientConfig, sigV4ClientConfig);



    apigClient.addfriendPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var addfriendPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/addfriend').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(addfriendPostRequest, authType, additionalParams, config.apiKey);
    };


    apigClient.checkincomingPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var checkincomingPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/checkincoming').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(checkincomingPostRequest, authType, additionalParams, config.apiKey);
    };


    apigClient.endcallPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var endcallPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/endcall').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(endcallPostRequest, authType, additionalParams, config.apiKey);
    };


    apigClient.findfriendPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var findfriendPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/findfriend').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(findfriendPostRequest, authType, additionalParams, config.apiKey);
    };


    apigClient.listfriendsPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var listfriendsPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/listfriends').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(listfriendsPostRequest, authType, additionalParams, config.apiKey);
    };


    apigClient.loginPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var loginPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/login').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(loginPostRequest, authType, additionalParams, config.apiKey);
    };


    apigClient.makecallPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var makecallPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/makecall').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(makecallPostRequest, authType, additionalParams, config.apiKey);
    };


    apigClient.signupPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var signupPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/signup').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(signupPostRequest, authType, additionalParams, config.apiKey);
    };


    apigClient.turnconfigPost = function (params, body, additionalParams) {
        if (additionalParams === undefined) { additionalParams = {}; }

        apiGateway.core.utils.assertParametersDefined(params, [], ['body']);

        var turnconfigPostRequest = {
            verb: 'post'.toUpperCase(),
            path: pathComponent + uritemplate('/turnconfig').expand(apiGateway.core.utils.parseParametersToObject(params, [])),
            headers: apiGateway.core.utils.parseParametersToObject(params, []),
            queryParams: apiGateway.core.utils.parseParametersToObject(params, []),
            body: body
        };


        return apiGatewayClient.makeRequest(turnconfigPostRequest, authType, additionalParams, config.apiKey);
    };


    return apigClient;
};
