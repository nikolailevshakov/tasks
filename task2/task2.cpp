#include <iostream>
#include <vector>
#include <fstream>
#include <string>
#include <sstream>
#include <algorithm>
#include <iomanip>


class Point
{
public:
    Point(double x_i, double y_i) :
        x(x_i), y(y_i) {}
    double x;
    double y;
};

std::vector<Point> get_points(std::string file_path) {

    std::ifstream input_file;
    input_file.open(file_path);
    if (!input_file) {
        std::cerr << "Can't open the input file" << std::endl;
        system("PAUSE");
        exit(1);
    }

    std::vector<Point> shape;
    std::string line;
    double x, y;
    while (std::getline(input_file, line))
    {
        std::istringstream ss(line);
        ss >> x >> y;
        shape.push_back(Point(x, y));
    }

    return shape;
}

std::vector<Point> extend_vector(std::vector<Point> vect)
{
    std::vector<Point> vect2(vect.begin(), vect.end());
    std::copy(vect.begin(), vect.end(), std::back_inserter(vect2));
    std::swap(vect, vect2);
    return vect;
}

bool in_range(double a, double b, double item)
{
    if (a > b && item > b && item < a)
        return true;
    if (a<b && item > a && item < b)
        return true;
    return false;
}

int check_place(std::vector<Point> shape_c, Point p_c)
{
    for (int i = 0; i < shape_c.size(); ++i)
    {
        if (p_c.x == shape_c[i].x && p_c.y == shape_c[i].y)
            return 0;

        if ((p_c.x == shape_c[i].x && p_c.x == shape_c[i + 1].x)
            &&
            (in_range(shape_c[i].y, shape_c[i+1].y, p_c.y)))
            return 1;

        if ((p_c.y == shape_c[i].y && p_c.y == shape_c[i + 1].y)
            &&
            (in_range(shape_c[i].x, shape_c[i+1].x, p_c.x)))
            return 1;

        if (shape_c[i].x == shape_c[i + 1].x)
            if (shape_c[i].y<p_c.y && shape_c[i + 1].y>p_c.y)
                return 2;

        if (shape_c[i].y == shape_c[i + 1].y)
            if (shape_c[i].x<p_c.x && shape_c[i + 1].x>p_c.x)
                return 2;

        if (i == shape_c.size() - 2)
            return 3;
    }
}



int main(int argc, char* argv[])
{
    argc = 2;
    std::vector<Point> shape;
    std::vector<Point> points;
    std::string shape_path = argv[1], points_path = argv[2];
    std::cout << argv[1] << " " << argv[2] << std::endl;

    shape = get_points(shape_path);
    points = get_points(points_path);

    std::vector<Point> ext_shape = extend_vector(shape);

    for (Point p : points)
    {
        std::cout << check_place(ext_shape, p) << std::endl;
    }

    system("PAUSE");
}

