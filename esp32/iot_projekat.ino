const int sensorPin = 27;

const int ledPins[] = {18, 19, 21, 22};
const int ledCount = 4;

void setup() {
  Serial.begin(115200);

  pinMode(sensorPin, INPUT);

  for (int i = 0; i < ledCount; i++) {
    pinMode(ledPins[i], OUTPUT);
    digitalWrite(ledPins[i], LOW);
  }
}

void loop() {
  int motion = digitalRead(sensorPin);

  if (motion == HIGH) {
    Serial.println("Motion detected");
    activateLeds();

    delay(2000);
  }
}

void activateLeds() {
  for (int j = 0; j < 5; j++) {        
    for (int i = 0; i < ledCount; i++) {
      digitalWrite(ledPins[i], HIGH);
      delay(100);
      digitalWrite(ledPins[i], LOW);
    }
  }
}
