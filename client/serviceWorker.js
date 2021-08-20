self.addEventListener("notificationclick", (event) => {
    const notification = event.notification;
    notification.close();
});

self.addEventListener("push", (event) => {
    const body = event.data.text();

    event.waitUntil(
        self.registration.showNotification("Notification", {
            body: body,
            vibrate: [100, 50, 100],
            requireInteraction: true,
        })
    );
});

const staticAssets = ["./", "./style.css", "./index.js"];

self.addEventListener("install", async (event) => {
    const cache = await caches.open("static-cache");
    cache.addAll(staticAssets);
});

self.addEventListener("fetch", (event) => {
    const req = event.request;
    const url = new URL(req.url);

    if (url.origin === location.url) {
        event.respondWith(cacheFirst(req));
    } else {
        event.respondWith(newtorkFirst(req));
    }
});

async function cacheFirst(req) {
    const cachedResponse = caches.match(req);
    return cachedResponse || fetch(req);
}

async function newtorkFirst(req) {
    const cache = await caches.open("dynamic-cache");

    try {
        const res = await fetch(req);
        cache.put(req, res.clone());
        return res;
    } catch (error) {
        return await cache.match(req);
    }
}
/*

self.addEventListener("push", (event) => {
    console.log("Push", event.data.text());
    
    const title = "Hello";
    const options = {
        body: event.data.text(),
    };

    event.waitUntil(self.registration.showNotification(title, options));
});

// TODO: Notification click event
self.addEventListener("notificationclick", function (event) {
    console.log("Push clicked");

    event.notification.close();

    event.waitUntil(clients.openWindow("https://github.com/boltlessengineer"));
});
*/
