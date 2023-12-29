import 'package:flame/game.dart';
import 'dart:ui';

class CozyGame extends FlameGame {
  @override
  Future<void> onLoad() async {
    // Initialize your game here
  }

  @override
  void update(double dt) {
    // Update your game state
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);
    final rect = Rect.fromLTWH(100, 100, 100, 100);
    final paint = Paint()..color = Color.fromARGB(255, 231, 53, 53);
    canvas.drawRect(rect, paint);
  }
}
