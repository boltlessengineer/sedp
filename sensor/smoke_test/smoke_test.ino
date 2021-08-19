/**
 * simple threshold test file
 */

#define pin A0
int amount = 0;

void setup() {
  // put your setup code here, to run once:
Serial.begin(9600);

}

void loop() {
  // put your main code here, to run repeatedly:
amount = analogRead(pin);
Serial.print("smoke: ");
Serial.println(amount);

}
