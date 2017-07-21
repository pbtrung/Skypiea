var options = {
    languageDictionary: {
        title: ""
    },
    theme: {
        logo: "/public/Skypiea.jpg"
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
$(".auth0-lock-header-logo").css({ "height": "85px", "margin-top": "5px" });