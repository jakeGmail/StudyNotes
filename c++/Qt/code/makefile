TARGET := QtStudy

INCLUDE := -I/usr/local/bin/Qt/include

LIB_DIR := -L/usr/local/bin/Qt/lib

LIB := \
 	-lQt5Widgets \
	-lQt5Core

SRCS :=
SRCS += main.cpp

OBJ := $(subst .cpp,.o,$(SRCS))

$(TARGET):$(OBJ)
	g++ $^ $(LIB_DIR) $(LIB) -o $@

%.o:%.cpp
	g++ -c $^ $(INCLUDE) -fPIC -DPIC -o $@

clean:
	-rm $(TARGET) $(OBJ)