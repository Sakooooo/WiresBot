{
  lib,
  buildGoModule,
  date,
  ...
}:
buildGoModule {
  pname = "wiresbot-go";
  version = "0-unstable-${date}";

  # Reproducible source path
  # pass a filter to get rid of irrelevant files
  # in the source path
  src = builtins.path {
    name = "wiresbot-go";
    path = ../.;
  };

  vendorHash = "sha256-nyB3juOxyrp+CuEpwKoBr23Py9S67FhUYg8sjLTBDPI=";
  ldflags = ["-s" "-w"];

  meta = {
    description = "Wired";
    homepage = "https://github.com/Sakooooo/WiresBot";
    license = lib.licenses.gpl3Only;
    mainProgram = "wiresbot";
  };
}
