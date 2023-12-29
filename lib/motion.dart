import 'package:vector_math/vector_math_64.dart';

/// A Motion can be used to define a Moveable entity with a Position.
class Motion {
  Vector2 start;
  int startUsec;
  // Speed of the entity in m/sec
  double speed;
  Vector2 dest;

  Motion(this.start, this.startUsec, this.speed, this.dest);

  factory Motion.idle(Vector2 pos, int currUsec) {
    return Motion(pos, currUsec, 0, pos);
  }

  Vector2 getStart() {
    return start;
  }

  Vector2 currentPosition(int currUsec) {
    if (start.distanceTo(dest) < 0.001) {
      return dest;
    }

    Vector2 distVec = dest - start;
    Vector2 dirVec = distVec.normalized();
    Vector2 speedVec = dirVec * speed;

    double deltaSec = (currUsec - startUsec) / 1000000.0;
    Vector2 currPos = start + speedVec * deltaSec;

    if (distVec.length2 < currPos.length2) {
      return dest;
    }
    return currPos;
  }

  Motion moveTo(int currUsec, double speed, Vector2 dest) {
    Vector2 currPos = currentPosition(currUsec);
    return Motion(currPos, currUsec, speed, dest);
  }
}
