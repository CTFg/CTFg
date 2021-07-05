import { ScoreboardTimelineQuery } from "../generated";
import {
  FlexibleXYPlot,
  HorizontalGridLines,
  LineSeries,
  LineSeriesPoint,
  XAxis,
  YAxis,
} from "react-vis";
import "../../node_modules/react-vis/dist/style.css";
import { useStyletron } from "baseui";

type ScoreboardTimelineProps = {
  scoreboard?: ScoreboardTimelineQuery["score_timeline"];
};


export default function ScoreboardTimeline({
  scoreboard,
}: ScoreboardTimelineProps) {
  const generateScoreboardTimeline = () => {
    const result: { [key: string]: Array<LineSeriesPoint> } = {};

    scoreboard?.map((item, idx) => {
      const timestampUTC: number = new Date(item?.event_time || 0).getTime();
      if (!result[item?.team_id || ""]) {
        result[item?.team_id || ""] = [
          { x: timestampUTC || 0, y: item?.score },
        ];
      } else {
        result[item?.team_id || ""].push({
          x: timestampUTC || 0,
          y: item?.score,
        });
      }
      return null;
    });
    return result;
  };

  const timelineData = generateScoreboardTimeline();
  const [css] = useStyletron();

  console.log(timelineData);

  return (
    <FlexibleXYPlot xType={"time"}>
      <XAxis />
      <YAxis />
      <HorizontalGridLines />
      {Object.keys(timelineData).map((team) => (
        <LineSeries
          curve={undefined}
          key={team}
          data={timelineData[team]}
          opacity={1}
          strokeStyle="solid"
          style={{}}
        />
      ))}
    </FlexibleXYPlot>
  );
}
