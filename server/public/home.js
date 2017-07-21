var options = {
    languageDictionary: {
        title: "Skypiea"
    },
    auth: {
        redirectUrl: AUTH0_CALLBACK_URL,
        responseType: "code",
        params: {
            scope: "openid profile"
        }
    }
};
var lock = new Auth0Lock(AUTH0_CLIENT_ID, AUTH0_DOMAIN, options);

lock.show();