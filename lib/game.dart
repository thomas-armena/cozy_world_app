import 'package:cozy_world_app/constants.dart';
import 'package:cozy_world_app/player.dart';
import 'package:flame/components.dart';
import 'package:flame/events.dart';
import 'package:flame/game.dart';
import 'dart:ui';

import 'package:flame/input.dart';

typedef MoveToCallback = void Function(Vector2 position);

class CozyGame extends FlameGame with TapDetector {
  late Player player;

  final int entityId;
  final MoveToCallback onMoveTo;

  CozyGame({required this.onMoveTo, required this.entityId});

  @override
  Future<void> onLoad() async {
    super.onLoad();

    // Create the player
    player = Player(Vector2(gridSize, gridSize), Vector2.all(gridSize));
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

  @override
  void onTapUp(TapUpInfo info) {
    // Handle tap up event
    Vector2 vector = info.eventPosition.global / gridSize;
    this.onMoveTo(vector);
  }
}
