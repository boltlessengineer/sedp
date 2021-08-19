/*
 * for testing connection with WIFI using WPA2_enterprise
 * if needed, combine this with main file, SEDP.ino
 */

extern "C" {
#include "user_interface.h"
#include "wpa2_enterprise.h"
}
#include <Arduino.h>
#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>

void delay_custom(int a) {
  long currentTime = millis();
  while (millis() - currentTime < a);
}

const char* c2s(wl_status_t status) {
  switch (status) {
    case WL_NO_SHIELD: return "WL_NO_SHIELD";
    case WL_IDLE_STATUS: return "WL_IDLE_STATUS";
    case WL_NO_SSID_AVAIL: return "WL_NO_SSID_AVAIL";
    case WL_SCAN_COMPLETED: return "WL_SCAN_COMPLETED";
    case WL_CONNECTED: return "WL_CONNECTED";
    case WL_CONNECT_FAILED: return "WL_CONNECT_FAILED";
    case WL_CONNECTION_LOST: return "WL_CONNECTION_LOST";
    case WL_DISCONNECTED: return "WL_DISCONNECTED";
  }
}

// SSID to connect to
static const char* ssid = "eduroam";
// Username for authentification
static const char* username = "*****";
// Password for authentification
static const char* password = "*****";
const int ledPin = 0;

void setup() {
  Serial.begin(115200);
  pinMode(ledPin, OUTPUT);

  Serial.print("Tryingonnect to ");
  Serial.println(ssid);

  wifi_station_disconnect();
  struct station_config wifi_config;

  memset(&wifi_config, 0, sizeof(wifi_config));
  strcpy((char*)wifi_config.ssid, ssid);
  strcpy((char*)wifi_config.password, password);
  wifi_station_set_config(&wifi_config);



  wifi_station_set_wpa2_enterprise_auth(1);
  wifi_station_set_enterprise_username((uint8*)username, strlen(username));
  wifi_station_set_enterprise_password((uint8*)password, strlen(password));
  wifi_station_connect();


  Serial.print("Status: ");
  Serial.println(wifi_station_get_connect_status());


  // Wait for connection AND IP address from DHCP
  while (WiFi.status() != WL_CONNECTED) {
    delay_custom(1000);
    ESP.wdtFeed();
    Serial.print(c2s(WiFi.status()));
    Serial.print(" IP: ");
    Serial.println(WiFi.localIP());
  }
} // setup

void loop() {

}
