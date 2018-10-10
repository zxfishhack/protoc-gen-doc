// AUTOGENERATED CODE. DO NOT EDIT.

package gendoc

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var embeddedResources = map[string]string{
	"docbook.tmpl": "H4sIAAAAAAAA/+xZ30/bOhR+719h5fEiEq64SFeTW6TBqmkChIDt3U1OW2uOndlOAUX93ycnIb8Tl7UUtvGCZJ/P58TH5/uOa/DpQ8jQCqSigo+df90jBwH3RUD5Yux8vZse/u+cTkaYSE19BpMRQlhTzWByLYUWvmDoXPhxCFwTTQXHXmYdIZQkkvAFIHdKGaj12ixV4BuUMReOksS9IiGs15W1ZnVEJEHuOShf0sisSl1U/F6CUmSRuy6dIxqMnSRxpzFjmWMnc1mNeCH4oiPqUFxjo3PkfiZqSoEFqpjHmswYoLkkIYwdwlgRsAiJfUaU4iRsRS8NKHPb+CDjYiFFHCFfMDV2/qs4RwibyQh8Y7yngV6OnX8cb2vEkXtiBx03IXoJJKjOIISluK/PIISBa/k4SXeLvWzQDbl7jGAYcUFmwIYhlZPsBGKv8Y3Ya20E65kIGusq9V2rBuvGKwU/9N2YUf4dmT/Ay4o2KTEVnVdRNsSegU2G/ZkVJlu2uJ0MyEr/HOYkZvobYbGJanCTfO4DSpKm3UsBSQI86Anayr1Jawqvn0c9+9jLGFHQ2ksJWFK46qEg7acHDdzI3O6JewVKQ4DKCBYOn+yDw2+D5UVOtmX6R6IsiKs4nIF8ZTHoqDJrjl5JENr+zgTXhHLKFw3PpeH5opMdy1+lOtirXXSqprJQeBzu7+qyK6VLk2yTt+N9yNuWumT29hvISZbvnUvJdrTcF6l6aJQP+tp6/UdGUd7mV8ghgxWw/j6NKZ8LGRI2yJb3Tv7eyd87+R/VyWu830R88hq5Bbmi/q89Qbx8D+/o35egl+JtvDG8uGBle0X2Tn8DP2JQGtml6wZUJLiCDaAvrk/5Ue5BnPL8NJQkn91aqp5y2nKfTT9bpt7ehaVhqEhJ54PorU8YkdlVOy20Ol3tl5T+K4qNhh32Extgt/Y2B1qnlJ+0G0mhRT8Tn+4SQpsE9gPODg6sTr6QFbGCrh/1UvA+WKPcWmxvc73sMmlB1Jnel5an1pM+z1foVBkP7cFohEmYFXUWRRt5M5nbCJhlrxfaYmuTqw2m1nna0eErlBxhr/j/xs8AAAD//2FyhLQRGQAA",
	"html.tmpl": "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",
	"markdown.tmpl": "H4sIAAAAAAAA/+xXW2/cRBR+7vyKw9oPiSpvxGu02QdaIoRKVdGKlygik2SyseQbtjci8liKeEkjVIJEWxAgLqKoXNQ0EQLUpJQ/0/GGf4HG4/XMeGe7Bom80Ke1fc6cOec737msBVlmd2+5qUfyHPUwBNgnS500jDoLfYRWVueseYQsC0ZfHLE/7qEscyDGwYBAd9n1SJLnKMvsLdcj7/KTsLgE3evY58Ze7cJKlo1tr85ZWVaJ5lGWVUbeIkmCB9zOJZP+8tDz5BkSbJb3VWdfD4a+cvBaGAyEbpuz76ckSNwwGBvg0Tge2SEeSFlpSEaX5w6pZQajN0m84260DMZOhLbETcpVfNLtcJNbNJnUbOR5V0O4dE76KH65Efb16ejbvdHJGfvqQ3b0+fMnB6tzVrKBPRw7O9gbEifdjUgyj6Qbs3JdE6f2QKVPLwLsuYNgqRO7g+200+9h2I7J1lLH4kTr3wqj3gLu9xaifsk1GSVwH/zNdQ+6V0myEbtR6oZBniMTgzQvJJg6kdtad7eg+wZOll3icfgRBfbo0+LoV6AgkAMKxTf7o0fP+MPh4fmfx0ARBcdxHDD/qHByq+DkOQKgUIMGVKR4N9JII97nheY1vE68UrX0PW74Ljy/Srbw0Evf4dnMc/jr7LPzowds7+kit6ALK2aU7ulckSio1YIoFAc/spN7KhKv4YQAhetDf53ELQExIiNvMqJTgyOer4RBit3ADQa8+FWJ8OSiYBKi3iuOAyTYBL+iJDhOH030q/Y0rYprGkkpsI/vjB4eAwV2eJ/tPZ0NvHSmDGsqyLPgU2CYDJ/wOMexmxikADCluTbBKO7vPz/7bXT6kN3+WbAPXSgPL5KF1esirE2QcG0qCyX6GtgNBsoBNWsA/eNW2oauxUffs8PvgML549+Lkw/qrLFP7rDTu+NX1MhO83FiONYkVsYjzMlZOGNUVgffJu8NSZKqvbf6pLdgoZtEYZCQWcpqkTQHulYF05x7AeBKgNMgvyTrr/KuGpYi/36Z9AkJv8jSUsTboe2rN9v+xN22bxyitv+fjVHbVwcprfGgL5ihVE5Qau5spppsMRjqXIOxJMeO1k1R4HJhg9X2J1paO7xWpjQ4oSklDYzHnc8M8r+dvjNAbvS9LCNewhltJLRe8AitCVHx5U/FwTN2+3jNeIVeLSKfapWJvmAss6boZZ29rLP/UZ3JkNrXmlYy47t0PQG7X05Vs0Y1W2tSVBPZsqrss8dP2Oldvp9EcZiGcit5sF/88gNQuHL5cv3xTbyD65cbu+l2GDT2lsbOYl5jtGd987xZ/hGXS426id3gDgoQO7DQB/1TtXOGKV/wqtU0ilQZ9159FwGMv9TY/B0AAP//c7ACVBUSAAA=",
	"scalars.json": "H4sIAAAAAAAA/9yXTWsaQRjH736KwVMKCYK1VnqrAcGz9hRSWHWybtnubnd3SqQUEpsSD6JQ1BQb2h5KWkqiPVj6EtN8mLiz5uRXKO6gO6Ozq4a0Id6W52Vm/r/5zwO7EQDgRQAAAIKarppquqDB4AMQzKkoI8PgKkkpqgmNYXgUyGoavzJr8OPieOUtWRXMaGSUeCI8F/gtWl5jesbxgplXFW5KR5nCKJFwEgEAXq56aGRa/SQyha5CJjwh8G6YJ5DV8a/1SQp1jLG+3tm5XftiVd9e1s/tbsP+uHuxU7Tav/qdD7j+jURwvWRX9vt/XvXOKoNu2Toq4veH1skBbn3vdZtkBVI/6Jb77R8kYjgbXuwUOQiZs7gIJYUDkKml8VHVFDxJMaEIdT4+qoWGF5dEBT0Fqg4S0vbwa0UwgA6fIUmHuTuzqLru/U9UoxFPqu5ZXKqyqohcrPxnR5dPcw0Zpi5RBZN4Q3T7NGNflmgBi3LkIw9XIa6t0DX7aobwK5kLLeAuDyA8QyC+I9DNWOJKZIwFrHL5pjR+erhdczrxuw55f8s+n4wFLDQNKhqZDepWj5wtaRvmOEayWs0xmUG3TMhYnz73flfCj8Mxu7mH6x0ymnG7hubx1M2Op2u2lcON4ytfbveiHG4zLXY7ppj/I5zHZks/iebxzHJOmYyqyvP8WdB1rkw6KnLXpFUO41BQOEInM4zCiSQtMK0juC4LhhFKCLJBPv1vmiHp3nP1q/V6t9/6+SidWIvhUuP+WkYywcPUejKJG/v48Jh/++xyLhc2LnpsT7NJsVec99iCJWOYeggpUlbNQR4dsiZYcUT5v4BMYUhiCovDpHd6iveq1smBfXy0GIZ4wYQpDxQbm8M9eSimu+bH4YfBuc21WDyZJiwCm4G/AQAA//+qprk/VRAAAA==",
}

func fetchResource(name string) ([]byte, error) {
	raw, ok := embeddedResources[name]
	if !ok {
		return nil, fmt.Errorf("Could not find resource for '%s'", name)
	}

	compressed, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	buf := bytes.NewBuffer(compressed)
	
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
