import 'dart:convert';

import 'package:cozy_world_app/protos/instance.pb.dart';
import 'package:flutter/material.dart';

int getEntityIdFromSnapshot(AsyncSnapshot<dynamic> snapshot) {
  if (!snapshot.hasData) return -1;
  return getEntityIdFromData(snapshot.data);
}

int getEntityIdFromData(List<int> data) {
  InstanceStreamResponse response = InstanceStreamResponse.fromBuffer(data);

  if (!response.hasConnectionCommand()) {
    return -1;
  }

  InstanceStreamResponse_ConnectionCommand connectionCommand =
      response.connectionCommand;
  // Extract the entityId and return it
  // This assumes that the JSON always contains a valid entityId field
  return connectionCommand.entityId;
}
