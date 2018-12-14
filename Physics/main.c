/*
 * Just some practice with the C syntax
 */

#include <assert.h>
#include <math.h>
#include <stdio.h>

double const GRAVITY_EARTH = 9.807;

double getAverageSpeed(double totalDistance, double totalTime) {
  return totalDistance / totalTime;
}

double getAcceleration(double initialVelocity, double finalVelocity,
                       double timeChange) {
  return (finalVelocity - initialVelocity) / timeChange;
}

// Newton's 2nd law
double getForce(double mass, double acceleration) {
  return mass * acceleration;
}

double getWeight(double mass) { return mass * GRAVITY_EARTH; }

double getSpringLoad(double springConstant, double springExtension) {
  return springConstant * springExtension;
}

double getKineticEnergy(double mass, double velocity) {
  return 0.5 * mass * pow(velocity, 2);
}

double getWork(double force, double distance) { return force * distance; }

double getWork2(double finalEnergy, double initialEnergy) {
  return finalEnergy - initialEnergy;
}

double getGravitationPotentialEnergy(double mass, double height) {
  return mass * height * GRAVITY_EARTH;
}

double calcAvg(double values[], int length) {
  int counter = 0;
  double sum = 0;

  for (int i = 0; i < length; i++) {
    counter++;
    sum += values[i];
  }
  return sum / counter;
}

int main() {
  assert(getAverageSpeed(100.0, 2.0) == 50.0);
  assert(getAcceleration(100.0, 0.0, 100.0) == -1.0);
  assert(getAcceleration(0.0, 100.0, 100.0) == 1.0);
  assert(getForce(3.0, 2.0) == 6.0);
  assert(getWeight(1.0) == GRAVITY_EARTH);
  assert(getSpringLoad(3.0, 2.0) == 6.0);
  assert(getKineticEnergy(2.0, 10.0) == 100.0);
  assert(getWork(3.0, 4.0) == 12.0);
  assert(getWork2(5.0, 1.0) == 4.0);
  assert(getGravitationPotentialEnergy(10.0, 100.0) == 9807.0);
  double testArray[] = {1.0, 2.0, 3.0};
  assert(calcAvg(testArray, sizeof(testArray) / sizeof(testArray[0])) == 2.0);
  return 0;
}
