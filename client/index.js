var serverKey =
    "BFUDTnUGRbMR9M8JU1zz-u5irMS6Z6uRF2aJSDNYweCtxCVF76eLsgnz10ca3PTDf9AH1M7-rQ-AZhgGIkIvz2o";

function URLB64ToUint8Array(base64String) {
    const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
    const base64 = (base64String + padding)
        .replace(/\-/g, "+")
        .replace(/_/g, "/");

    const rawData = window.atob(base64);
    const outputArray = new Uint8Array(rawData.length);

    for (let i = 0; i < rawData.length; ++i) {
        outputArray[i] = rawData.charCodeAt(i);
    }
    return outputArray;
}

function requestTo(url, jsonString) {
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(jsonString);
}

if ("serviceWorker" in navigator && "Notification" in window) {
    if (serverKey === undefined || serverKey == null) {
        console.error("Missing ServerKey!");
        //return;
    }
    navigator.serviceWorker.register("serviceWorker.js").then(() => {
        console.log("Service worker registered");

        navigator.serviceWorker.ready.then((registration) => {
            registration.pushManager.getSubscription().then((subscription) => {
                if (subscription === undefined || subscription == null) {
                    navigator.serviceWorker
                        .getRegistration()
                        .then((registration) => {
                            registration.pushManager
                                .subscribe({
                                    userVisibleOnly: true,
                                    applicationServerKey:
                                        URLB64ToUint8Array(serverKey), //This is a custom function for convert key in bytes array
                                })
                                .then((subscription) => {
                                    const json = JSON.stringify(
                                        subscription.toJSON(),
                                        null,
                                        2
                                    );

                                    // Add Login to update subscription info to your Server
                                    requestTo("/sub", json);
                                });
                        });
                } else {
                    const json = JSON.stringify(subscription.toJSON(), null, 2);

                    // Add Login to update subscription info to your Server
                    requestTo("/sub", json);
                }
            });
        });
    });
}
