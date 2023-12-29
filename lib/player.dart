import 'dart:ui';

import 'package:flame/components.dart';

class Player extends PositionComponent {
  Player(Vector2 position, Vector2 size)
      : super(position: position, size: size);

  @override
  void render(Canvas canvas) {
    super.render(canvas);
    final paint = Paint()..color = Color.fromARGB(255, 254, 92, 92);
    canvas.drawRect(size.toRect(), paint);
  }
}
