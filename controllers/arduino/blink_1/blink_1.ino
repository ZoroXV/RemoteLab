const int blueLedPin = 2;
const int yellowLedPin = 3;
const int whiteLedPin = 4;
const int redLedPin = 5;

void setup() {
  // initialize digital pin LED_BUILTIN as an output.
  pinMode(blueLedPin, OUTPUT);
  pinMode(yellowLedPin, OUTPUT);
  pinMode(whiteLedPin, OUTPUT);
  pinMode(redLedPin, OUTPUT);
}

// the loop function runs over and over again forever
void loop() {
  digitalWrite(blueLedPin, HIGH);
  
  delay(100);
  
  digitalWrite(blueLedPin, LOW);
  digitalWrite(whiteLedPin, HIGH);
  
  delay(100);
  
  digitalWrite(whiteLedPin, LOW);
  digitalWrite(redLedPin, HIGH);
  
  delay(100);
  
  digitalWrite(redLedPin, LOW);
}
