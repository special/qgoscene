#ifndef SCENE_H
#define SCENE_H

#ifdef __cplusplus
extern "C" {
#endif

void createApplication(int argc, char **argv);
void createEngine();
void loadScene(char *path);
void loadSceneData(char *data);
void addImportPath(char *path);
void setImportPathList(int n, char **paths);
void setContextProperty(char *name, char *value);
int execApplication();
void quitApplication();

#ifdef __cplusplus
}
#endif

#endif
