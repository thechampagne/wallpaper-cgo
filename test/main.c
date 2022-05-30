#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <limits.h>
#include "wallpaper.h"

int main()
{
  wallpaper background = wallpaper_get();
  if (background.error != 0)
  {
    printf("%s\n", background.error);
    free(background.error);
    free(background.buffer);
  } else {
    printf("Current wallpaper: %s\n", background.buffer);
    free(background.error);
    free(background.buffer);
  }

  char* set_from_url = wallpaper_set_from_url("https://source.unsplash.com/1920x1080");
  if (set_from_url != 0)
  {
    printf("set_from_url error: %s\n", set_from_url);
  }
  free(set_from_url);

  sleep(5);

  char* cwd = (char*) malloc(PATH_MAX + 14);
  getcwd(cwd, PATH_MAX + 14);
  strcat(cwd, "/test/test.jpg");
  char* set_from_file = wallpaper_set_from_file(cwd);
  if (set_from_file != 0)
  {
    printf("set_from_file error: %s\n", set_from_file);
  }
  free(set_from_file);
  free(cwd);

  /*
  typedef enum {
    Center,
    Crop,
    Fit,
    Span,
    Stretch,
    Tile
  } wallpaper_mode;
  */
  char* set_mode = wallpaper_set_mode(Center);
  if (set_mode != 0)
  {
    printf("set_mode error: %s\n", set_mode);
  }
  free(set_mode);
}