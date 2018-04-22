import QtQuick 2.9
import QtQuick.Window 2.2

Window {
    width: 600
    height: 600
    visible: true

    Text {
        anchors.centerIn: parent
        font.pixelSize: 26
        // location is provided by SetContextProperty
        text: "Hello " + location + "!"
    }
}
