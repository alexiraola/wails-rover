import { Events } from "@wailsio/runtime";
import { useEffect, useState } from "preact/hooks";
import { LocationParser } from "./location.parser";
import { Orientation, Position } from "./location";
import { Config } from '../bindings/wails_rover';

export default function useLocation() {
  const [orientation, setOrientation] = useState(Orientation.NORTH);
  const [position, setPosition] = useState(new Position(0, 0));
  const [worldSize, setWorldSize] = useState([10, 10]);

  useEffect(() => {
    const unregister = Events.On('location', (message: any) => {
      const location = LocationParser.parse(message.data);
      setOrientation(location.orientation);
      setPosition(location.position);
    });

    return () => unregister();
  }, []);

  useEffect(() => {
    Config.WorldSize().then(size => {
      setWorldSize(size);
    });
  }, []);

  return { orientation, position, worldSize }
}
