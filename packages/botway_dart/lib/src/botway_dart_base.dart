import 'dart:io';
import 'dart:convert';
import 'package:yaml/yaml.dart';
import 'package:path/path.dart' as p;

class Botway {
  get_bot_info(String value) {
    if (File(".botway.yaml").existsSync()) {
      String BotConfigFile = File('.botway.yaml').readAsStringSync();
      dynamic BotConfig = loadYaml(BotConfigFile);

      return BotConfig["bot"][value];
    } else {
      throw Exception("ERROR: Bot config file not found");
    }
  }

  String? get HomePath =>
      Platform.environment['HOME'] ?? Platform.environment['USERPROFILE'];

  String BotwayConfigFile() {
    File BWFile = File(p.join(HomePath.toString(), '.botway', 'botway.json'));

    if (BWFile.existsSync()) {
      return BWFile.readAsStringSync();
    } else {
      throw Exception("ERROR: Botway config file not found");
    }
  }

  dynamic get BotwayConfig => json.decode(BotwayConfigFile());

  Get_Token() {
    return BotwayConfig["botway"]["bots"][get_bot_info("name")]["bot_token"];
  }

  Get_App_Id() {
    String value = "bot_app_id";

    if (get_bot_info("type") == "slack") {
      value = "bot_app_token";
    }

    return BotwayConfig["botway"]["bots"][get_bot_info("name")][value];
  }

  Get_Guild_Id(String server_name) {
    if (get_bot_info("type") != "discord") {
      throw Exception("ERROR: This function/feature is only working with discord bots");
    } else {
      return BotwayConfig["botway"]["bots"][get_bot_info("name")]["guilds"][server_name]["server_id"];
    }
  }

  Get_Signing_Secret() {
    if (get_bot_info("type") != "slack") {
      throw Exception("ERROR: This function/feature is only working with slack bots");
    } else {
      return BotwayConfig["botway"]["bots"][get_bot_info("name")]["signing_secret"];
    }
  }
}
