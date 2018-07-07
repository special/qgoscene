#include "scene.h"

#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include <QQmlContext>

static QGuiApplication *app = nullptr;
static QQmlApplicationEngine *engine = nullptr;

void createApplication(int tmpargc, char **argv)
{
    if (app)
        return;

    // argv from Go, and all of the strings in it, are freshly malloc'd
    // and nothing will free them, so those are safe to pass to Qt as-is.
    // The constructor takes and stores an int& for argc, so we need a
    // copy for that with more permanent storage.
    static int argc = tmpargc;

    app = new QGuiApplication(argc, argv);
}

void createEngine()
{
    if (engine)
        return;
    engine = new QQmlApplicationEngine(app);
}

void loadScene(char *path)
{
    engine->load(QString::fromUtf8(path));
    free(path);
}

void loadSceneData(char *data)
{
    QByteArray d(data);
    engine->loadData(d);
    free(data);
}

void addImportPath(char *path)
{
    engine->addImportPath(QString::fromUtf8(path));
    free(path);
}

void setImportPathList(int n, char **paths)
{
    QStringList p;
    for (int i = 0; i < n; i++) {
        p << QString::fromUtf8(paths[i]);
        free(paths[i]);
    }
    free(paths);

    engine->setImportPathList(p);
}

void setContextProperty(char *name, char *value)
{
    engine->rootContext()->setContextProperty(QString::fromUtf8(name), QString::fromUtf8(value));
    free(name);
    free(value);
}

int execApplication()
{
    if (!app || !engine)
        return 1;
    return app->exec();
}

void quitApplication()
{
    app->quit();
}
