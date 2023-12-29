import 'dart:convert';

import 'package:flutter/material.dart';

int getEntityIdFromSnapshot(AsyncSnapshot<dynamic> snapshot) {
  if (!snapshot.hasData) return -1;
  return getEntityIdFromData(snapshot.data);
}

int getEntityIdFromData(String jsonString) {
  // Decode the JSON string into a Map
  Map<String, dynamic> data = jsonDecode(jsonString);

  // Extract the entityId and return it
  // This assumes that the JSON always contains a valid entityId field
  return data['entityId'];
}
