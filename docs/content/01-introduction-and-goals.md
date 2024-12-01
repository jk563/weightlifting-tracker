# Introduction and Goals

The Weightlifting Tracker records successful workouts, and computes the next workout to be completed based on a steady progression to help the user incrementally gain strength across a number of exercises. The primary goal is to have a dynamic progressive weightlifting plan without having to remember previous workouts and work out the next one to do.

## Requirements Overview
| ID    | Name                  | Description |
| :---: | ---------             | ------------------------ |
| R1    | Get next workouts     | View the next workout that is due to be completed, based on incremental progression and lowering weight when there hasn't been a recent success |
| R2    | Record workout result | Store the success of a workout |

## Quality Goals
| ID    | Quality     | Description |
| :---: | -----       | --------------------- |
| Q1    | Correctness | The correct workout details will be returned, following an incremental progression with weight removed as the time between successful workouts increases |
| Q2    | Cost        | The solution should cost a minimal amount to develop and run, as it does not generate any revenue |
| Q3    | Clarity     | The system may go long periods of time without being worked on, it should be easy to come in fresh and understand it |

## Stakeholders
| Role       | Name        | Expectations |
| ---        | ---         | ---------    |
| Everything | Jamie Kelly | The project is cheap, produces the right output, and is easy to understand |

