import { Orientation, Position, Location } from "./location";

export class LocationParser {
  static parse(location: string): Location {
    const parts = location.split(':');
    const orientation = this.parseOrientation(parts[0]);
    const position = this.parsePosition(parts[1], parts[2]);

    return new Location(orientation, position);
  }

  private static parseOrientation(orientation: string): Orientation {
    switch (orientation) {
      case 'N':
        return Orientation.NORTH;
      case 'E':
        return Orientation.EAST;
      case 'S':
        return Orientation.SOUTH;
      case 'W':
        return Orientation.WEST;
    }
    throw new Error(`Invalid orientation: ${orientation}`);
  }

  private static parsePosition(x: string, y: string): Position {
    return new Position(this.parseIntOrThrow(x), this.parseIntOrThrow(y));
  }

  private static parseIntOrThrow(s: string): number {
    const integerRegex = /^-?\d+$/;

    if (!integerRegex.test(s)) {
      throw new Error(`Invalid position: ${s}`);
    }

    const number = parseInt(s, 10);
    if (isNaN(number)) {
      throw new Error(`Invalid position: ${s}`);
    }
    return number;
  }
}

