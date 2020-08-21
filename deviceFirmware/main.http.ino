#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <ArduinoJson.h>

/* This is the basic setup to have a device connect to WIFI and 
  send a post to the specified server with a JSON body that contains the ID and sensor reading

  For this to be truly usable the WIFI settings need to be configurable after flashing. 
  i.e. start wifi-host allow device to connect, show config http page, save values to EEPROM, 
  switch to WIFI client, try to connect, if connection fails return to host mode so user can fix settings.



All the serial stuff can be removed but its nice for testing
*/

#define SERVER_IP "192.168.1.230:8282"

#ifndef STASSID
#define STASSID "<WIFI_SSID>"
#define STAPSK  "<WIFI_PASSWORD>"
#define DEVICEID "A141414" // this needs to be unique for each device
#define REPORTDELAY 10000  // time in milisec between status reports
#endif

const int analogInPin = A0;  // ESP8266 Analog Pin ADC0 = A0
int sensorValue = 0;  // value read from the pot

void setup() {

  Serial.begin(115200);

  Serial.println();
  Serial.println();
  Serial.println();

  WiFi.begin(STASSID, STAPSK);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected! IP address: ");
  Serial.println(WiFi.localIP());

  

}

void loop() {

  // read the analog in value
  sensorValue = analogRead(analogInPin);
  // wait for WiFi connection
  if ((WiFi.status() == WL_CONNECTED)) {
    String json;
    const int capacity = JSON_ARRAY_SIZE(2);
    StaticJsonDocument<capacity> doc;
    doc["value"] = sensorValue;
    doc["id"] = DEVICEID;

    WiFiClient client;
    HTTPClient http;

    serializeJson(doc, json);

    Serial.print("[HTTP] begin...\n");
    // configure traged server and url
    http.begin(client, "http://" SERVER_IP "/report"); //HTTP
    http.addHeader("Content-Type", "application/json");

    Serial.print("[HTTP] POST...\n");
    // start connection and send HTTP header and body
    int httpCode = http.POST(json);

    // httpCode will be negative on error
    if (httpCode > 0) {
      // HTTP header has been send and Server response header has been handled
      Serial.printf("[HTTP] POST... code: %d\n", httpCode);

      // file found at server
      if (httpCode == HTTP_CODE_OK) {
        const String& payload = http.getString();
        Serial.println("received payload:\n<<");
        Serial.println(payload);
        Serial.println(">>");
      }
    } else {
      Serial.printf("[HTTP] POST... failed, error: %s\n", http.errorToString(httpCode).c_str());
    }

    http.end();
  }

  delay(REPORTDELAY);
}