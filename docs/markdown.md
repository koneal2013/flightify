# Flightify API

Api documentation for Flightify service.

## Informations

### Version

1.0

### License

[Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### Contact

Kenston O'Neal koneal2013@gmail.com

## Content negotiation

### URI Schemes

* http

### Consumes

* application/json

### Produces

* application/json

## All endpoints

### operations


| Method | URI        | Name                              | Summary |
| ------ | ---------- | --------------------------------- | ------- |
| GET    | /status    | [get status](#get-status)         |         |
| POST   | /calculate | [post calculate](#post-calculate) |         |

## Paths

### <span id="get-status"></span> get status (*GetStatus*)

```
GET /status
```

Return 200 OK if server is ready to accept requests

#### All responses


| Code                   | Status | Description | Has headers | Schema                           |
| ---------------------- | ------ | ----------- | :---------: | -------------------------------- |
| [200](#get-status-200) | OK     | OK          |            | [schema](#get-status-200-schema) |

#### Responses

##### <span id="get-status-200"></span> 200 - OK

Status: OK

###### <span id="get-status-200-schema"></span> Schema

### <span id="post-calculate"></span> post calculate (*PostCalculate*)

```
POST /calculate
```

Returns a flight itinerary with origin airport code and final destination airport code

#### Consumes

* application/json

#### Produces

* application/json

#### Parameters


| Name  | Source | Type   | Go type  | Separator | Required | Default | Description                                                      |
| ----- | ------ | ------ | -------- | --------- | :------: | ------- | ---------------------------------------------------------------- |
| input | `body` | string | `string` |           |    ✓    |         | slice of flight segments (i.e. [['ATL', 'EWR'], ['SFO', 'ATL']]) |

#### All responses


| Code                       | Status      | Description | Has headers | Schema                               |
| -------------------------- | ----------- | ----------- | :---------: | ------------------------------------ |
| [200](#post-calculate-200) | OK          | OK          |            | [schema](#post-calculate-200-schema) |
| [400](#post-calculate-400) | Bad Request | Bad Request |            | [schema](#post-calculate-400-schema) |

#### Responses

##### <span id="post-calculate-200"></span> 200 - OK

Status: OK

###### <span id="post-calculate-200-schema"></span> Schema

##### <span id="post-calculate-400"></span> 400 - Bad Request

Status: Bad Request

###### <span id="post-calculate-400-schema"></span> Schema

## Models
