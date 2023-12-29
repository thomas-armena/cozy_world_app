import 'package:cozy_world_app/game.dart';
import 'package:cozy_world_app/protos/instance.pb.dart';
import 'package:cozy_world_app/protos/math.pb.dart';
import 'package:cozy_world_app/utils/getEntityIdFromData.dart';
import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    const title = 'WebSocket Demo';
    return const MaterialApp(
      title: title,
      home: MyHomePage(
        title: title,
      ),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({
    super.key,
    required this.title,
  });

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final TextEditingController _controller = TextEditingController();
  final _channel = WebSocketChannel.connect(
    Uri.parse('ws://10.88.111.18:8080'),
  );

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
          body: StreamBuilder(
        stream: _channel.stream,
        builder: (context, snapshot) {
          int entityId = getEntityIdFromSnapshot(snapshot);
          print(entityId);
          return GameWidget(
              game: CozyGame(entityId: entityId, onMoveTo: this._moveTo));
        },
      )),
    );
  }

  void _moveTo(Vector2 vector) {
    var obj = InstanceStreamRequest.create();
    obj.moveToCommand = InstanceStreamRequest_MoveToCommand(
        position: Vec2(x: vector.x, y: vector.y));
    var data = obj.writeToBuffer();
    print(data);
    _channel.sink.add(data);
  }

  void _sendMessage() {
    if (_controller.text.isNotEmpty) {
      _channel.sink.add(_controller.text);
    }
  }

  @override
  void dispose() {
    _channel.sink.close();
    _controller.dispose();
    super.dispose();
  }
}
