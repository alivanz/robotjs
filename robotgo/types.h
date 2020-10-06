#include <stdint.h>

typedef struct {
	uint8_t    Kind;
	uint64_t   When;
	uint16_t   Mask;
	uint16_t   Reserved;

	uint16_t   Keycode;
	uint16_t   Rawcode;
	uint8_t    Keychar;

	uint16_t   Button;
	uint16_t   Clicks;

	int16_t    X;
	int16_t    Y;

	uint16_t   Amount;
	int32_t    Rotation;
	uint8_t    Direction;
} Event_t;

typedef void (*fevent_hook)(Event_t*);
