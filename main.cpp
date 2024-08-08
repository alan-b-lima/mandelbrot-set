#include "../Win32API.h"
#include <complex>

using namespace Win32API;
using complex = std::complex<float>;

struct {
   float originX, originY, scale;
   int max_interations;

   float lower_bound, upper_bound;
   bool update;
} set;

int main(int argc, char** argv) {

   set.max_interations = atoi(argv[1]);
   set.scale = atof(argv[2]);

   set.lower_bound = atof(argv[3]);
   set.upper_bound = atof(argv[4]);
   
   set.update = true;

   if (set.scale == 0) return 1;

   Window& window = window.getInstance();
   window.launchWindow(SW_SHOWMAXIMIZED);

   return 0;
}

void Window::update(float) {
   if (set.update) {

      for (int y = 0; y < ctx.height(); y++) {
         for (int x = 0; x < ctx.width(); x++) {
            float x_ = (float)(x - set.originX) / set.scale;
            float y_ = (float)(y - set.originY) / set.scale;
            float real = 0.f;
            float imag = 0.f;

            for (int t = 0; t < set.max_interations; t++) {
               float helper = real * real - imag * imag + x_;
               imag = 2 * real * imag + y_;
               real = helper;

               helper = real * real + imag * imag;
               if (set.lower_bound <= helper && helper <= set.upper_bound) {

                  helper = (float)t * 255.f / set.max_interations;

                  ctx.fillPixel(x, y, (clr)clamp(0.f, helper, 255.f) << 16
                     //rgb(
                     //   clamp(0.f, helper, 1.f),
                     //   clamp(0.f, (t * imag) * .25, 1.f),
                     //   clamp(0.f, (t * real) * .25, 1.f))
                  );
               }
            }
         }
      }

      set.update = false;
   }

   if (input.keyboard.pressed(VK_ESCAPE)) exit(0);
}

void Window::resize() {
   set.originX = (float)(ctx.width() >> 1);
   set.originY = (float)(ctx.height() >> 1);

   set.update = true;
}