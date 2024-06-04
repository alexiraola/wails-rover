import car from './assets/car.svg';
import { useMemo } from 'preact/hooks';
import { Orientation, Position } from './location';

export type RoverProps = {
  orientation: Orientation,
  position: Position
}

const CELL_SIZE = 50;

export function Rover({ orientation, position }: RoverProps) {
  const rotation = useMemo(() => {
    switch (orientation) {
      case Orientation.NORTH:
        return '-90deg';
      case Orientation.EAST:
        return '0deg';
      case Orientation.SOUTH:
        return '90deg';
      case Orientation.WEST:
        return '180deg';
    }
  }, [orientation]);

  const translation = useMemo(() => {
    return `${position.x * CELL_SIZE}px, ${position.y * -CELL_SIZE}px`;
  }, [position]);

  return <div style={{ transform: `translate(${translation})` }} class="rover"><img style={{ transform: `rotate(${rotation})` }} src={car} /></div>
}
