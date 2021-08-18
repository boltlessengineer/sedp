let appServerPublicKey =
    "BKKYv6LblA1kHsUtXHDmhxCwXMPp-AMwdTavbuYoMKnsWNXth9SHCsdBsPWC1COmRfWzb7CUErNs-HPGkwyY7wk";
let isSubcribed = false;
let swRegist = null;

function urlB64ToUint8Array(base64String) {
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

function updateButton() {
    const pushBtn = document.getElementById("btn");
    if (isSubcribed) {
        pushBtn.textContent = "Disable Push Messaging";
    } else {
        pushBtn.textContent = "Enable Push Messaging";
    }
    pushBtn.disabled = false;
}

function updateSubscription(subscription) {
    let detailArea = document.getElementById("subscription_detail");

    if (subscription) {
        detailArea.innerText = JSON.stringify(subscription);
        detailArea.parentElement.classList.remove("hide");
    } else {
        detailArea.parentElement.classList.add("hide");
    }
}

function subscribe() {
    const applicationServerKey = urlB64ToUint8Array(appServerPublicKey);
    swRegist.pushManager
        .subscribe({
            userVisibleOnly: true,
            applicationServerKey: applicationServerKey,
        })
        .then((subscription) => {
            console.log("User is subscribed.");
            updateSubscription(subscription);
            isSubcribed = true;
            updateButton();
        })
        .catch((err) => {
            console.log("Failed to subscribe the user: ", err);
            updateButton();
        });
}

function unsubscribe() {
    swRegist.pushManager
        .getSubscription()
        .then((subscription) => {
            if (subscription) {
                return subscription.unsubscribe();
            }
        })
        .catch((error) => {
            console.log("Error unsubscribing", error);
        })
        .then(() => {
            updateSubscription(null);
            console.log("User is unsubscribed.");
            isSubscribed = false;
            updateButton();
        });
}

function initPush() {
    const pushBtn = document.getElementById("btn");
    pushBtn.addEventListener("click", () => {
        if (isSubcribed) {
            // todo
        } else {
            subscribe();
        }
    });
    swRegist.pushManager.getSubscription().then(function (subscription) {
        isSubcribed = !(subscription === null);
        updateSubscription(subscription);

        if (isSubcribed) {
            console.log("User is subscribed.");
        } else {
            console.log("User is not subscribed");
        }

        updateButton();
    });
}

if ("serviceWorker" in navigator) {
    console.log("Service Worker and Push is supported");

    navigator.serviceWorker
        .register("serviceWorker.js")
        .then(function (regist) {
            console.log("Service Worker is registered", regist);

            swRegist = regist;

            initPush();
        })
        .catch(function (error) {
            console.error("Service Worker Error", error);
        });
} else {
    console.warn("Push messaging is not supported");
    pushButton.textContent = "Push Not Supported";
}
