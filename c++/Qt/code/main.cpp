#include <iostream>
#include <QtWidgets/QApplication>
#include <QtWidgets/QLabel>


int main(int argc ,char *argv[])
{
    QApplication a(argc, argv);
    QLabel label("hello,world!");
    label.resize(200,200);
    label.show();
    return a.exec();
}

