import 'package:cozy_world_app/constants.dart';
import 'package:cozy_world_app/player.dart';
import 'package:flame/components.dart';
import 'package:flame/game.dart';
import 'dart:ui';

class CozyGame extends FlameGame {
  late Player player;

  int entityId;

  CozyGame(this.entityId);

  @override
  Future<void> onLoad() async {
    super.onLoad();

    // Create the player
    player = Player(Vector2(100, 100), Vector2.all(gridSize));
    add(player);

    // Make the camera follow the player
    camera.follow(player);
  }

  @override
  void update(double dt) {
    super.update(dt);
    // Update your game state
  }

  void drawGrid(Canvas canvas) {
    final paint = Paint()
      ..color = Color(0xFF00FF00)
      ..style = PaintingStyle.stroke;

    for (double x = 0; x < size.x; x += gridSize) {
      for (double y = 0; y < size.y; y += gridSize) {
        canvas.drawRect(Rect.fromLTWH(x, y, gridSize, gridSize), paint);
      }
    }
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);
    // Render the grid
    drawGrid(canvas);
  }
}
