import car from './assets/car.svg';
import { useMemo } from 'preact/hooks';

export type RoverProps = {
  location: string
}

const CELL_SIZE = 50;

export function Rover({ location }: RoverProps) {
  const rotation = useMemo(() => {
    switch (location.substring(0, 1)) {
      case 'N':
        return '-90deg';
      case 'E':
        return '0deg';
      case 'S':
        return '90deg';
      case 'W':
        return '180deg';
    }
    return '0deg';
  }, [location]);

  const translation = useMemo(() => {
    const parts = location.split(':');

    return `${parseInt(parts[1]) * CELL_SIZE}px, ${parseInt(parts[2]) * -CELL_SIZE}px`;
  }, [location]);

  console.log(rotation, translation);

  return <div style={{ transform: `translate(${translation})` }} class="rover"><img style={{ transform: `rotate(${rotation})` }} src={car} /></div>
}
