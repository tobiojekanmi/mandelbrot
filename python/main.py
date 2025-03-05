from typing import List


def calculate_mandelbrot(
    max_iters: int,
    x_min: float,
    x_max: float,
    y_min: float,
    y_max: float,
    width: int,
    height: int,
) -> List[List[int]]:

    rows = []
    for img_y in range(height):
        row = []
        for img_x in range(width):
            x_percent = img_x / width
            y_percent = img_y / height
            cx = x_min + (x_max - x_min) * x_percent
            cy = y_min + (y_max - y_min) * y_percent
            escaped_at = mandelbrot_at_point(cx, cy, max_iters)
            row.append(escaped_at)
        rows.append(row)

    return rows


def mandelbrot_at_point(cx: float, cy: float, max_iters: int) -> int:
    z = complex(0.0, 0.0)
    c = complex(cx, cy)

    for i in range(max_iters):
        if abs(z) > 2.0:
            return i
        z = z * z + c
    return max_iters


def render_mandelbrot(escape_vals: List[List[int]]) -> None:
    print("")
    for row in escape_vals:
        line = ""
        for val in row:
            if val >= 0 and val <= 2:
                line += " "
            elif val >= 3 and val <= 5:
                line += "."
            elif val >= 6 and val <= 10:
                line += "•"
            elif val >= 11 and val <= 30:
                line += "*"
            elif val >= 31 and val <= 100:
                line += "+"
            elif val >= 101 and val <= 200:
                line += "x"
            elif val >= 201 and val <= 400:
                line += "$"
            elif val >= 401 and val <= 700:
                line += "#"
            else:
                line += "%"
        print(line)
    print("")


def main():
    mandelbrot = calculate_mandelbrot(1000, -2.0, 1.0, -1.0, 1.0, 100, 24)
    render_mandelbrot(mandelbrot)


if __name__ == "__main__":
    main()
