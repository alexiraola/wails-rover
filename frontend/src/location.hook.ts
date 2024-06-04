import { Events } from "@wailsio/runtime";
import { useEffect, useState } from "preact/hooks";
import { LocationParser } from "./location.parser";
import { Orientation, Position } from "./location";

export default function useLocation() {
  const [orientation, setOrientation] = useState(Orientation.NORTH);
  const [position, setPosition] = useState(new Position(0, 0));

  useEffect(() => {
    const unregister = Events.On('location', (message: any) => {
      const location = LocationParser.parse(message.data);
      setOrientation(location.orientation);
      setPosition(location.position);
    });

    return () => unregister();
  }, []);

  return { orientation, position }
}
