#include <assert.h>
#include <stdio.h>

double const GRAVITY_EARTH = 9.807;

double getAverageSpeed(double totalDistance, double totalTime)
{
    return totalDistance / totalTime;
}

double getAcceleration(double initialVelocity, double finalVelocity,
                       double timeChange)
{
    return (finalVelocity - initialVelocity) / timeChange;
}

// Newton's 2nd law
double getForce(double mass, double acceleration)
{
    return mass * acceleration;
}

double getWeight(double mass)
{
    return mass * GRAVITY_EARTH;
}

int main()
{
    assert(getAverageSpeed(100.0, 2.0) == 50.0);
    assert(getAcceleration(100.0, 0.0, 100.0) == -1.0);
    assert(getAcceleration(0.0, 100.0, 100.0) == 1.0);
    assert(getForce(3.0, 2.0) == 6.0);
    assert(getWeight(1.0) == GRAVITY_EARTH);
    return 0;
}