export enum Orientation {
  NORTH = 'N',
  EAST = 'E',
  SOUTH = 'S',
  WEST = 'W'
}

export class Position {
  constructor(readonly x: number, readonly y: number) { }
}


export class Location {
  constructor(readonly orientation: Orientation, readonly position: Position) { }
}

