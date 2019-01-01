package rank

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"gitlab.com/oiacow/nextfesl/network"
)


type ansUpdateStats struct {
	Txn   string      `fesl:"TXN"`
	Users []userStats `fesl:"u"`
}

type userStats struct {
	OwnerID   int          `fesl:"o"`  // 3
	OwnerType int          `fesl:"ot"` // 1
	Stats     []updateStat `fesl:"s"`
}

type updateStat struct {
	Key        string `fesl:"k"`  // c_ltp
	PointType  int    `fesl:"pt"` // 0
	Text       string `fesl:"t"`  // ""
	UpdateType int    `fesl:"ut"` // 0
	Value      string `fesl:"v"`  // 9025.0000
}

type stat struct {
	val float64
	stringval  string //omg wow
	floatval float64
	//textval string
}

// UpdateStats - updates stats about a soldier
func (r *Ranking) UpdateStats(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case "server":
		r.serverUpdateStats(&event)
	default:
		r.clientUpdateStats(&event)
	}
}

func (r *Ranking) clientUpdateStats(event *network.EventClientCommand) {
	r.updateStats(event)
}

func (r *Ranking) serverUpdateStats(event *network.EventClientCommand) {
	r.updateStats(event)
}

func (r *Ranking) updateStats(event *network.EventClientCommand) {
	reply := event.Command.Message
	users, _ := strconv.Atoi(event.Command.Message["u.[]"])
	sess := r.DB.NewSession()

	for i := 0; i < users; i++ {
		heroID, _ := event.Command.Message.IntVal(fmt.Sprintf("u.%d.o", i))
		p, err := r.DB.FindHeroStats(sess, heroID)
		if err != nil {
			logrus.
				WithError(err).
				WithField("heroID", event.Command.Message[fmt.Sprintf("u.%d.o", i)]).
				Warn("Cant Query hero stats"+ heroID)
			return
		}

		// keys := event.Command.Message.ArrayStrings("keys")
		// stats, err := r.fetchStats(&heroStats, keys)
		// if err != nil {
		// 	logrus.WithError(err).Error("Cannot retrieve given stats")
		// 	return		}	
		
		// for i := 0; i < count; i++ {
		// arr[i] = m.Get(fmt.Sprintf("%s.%d", prefix, i))
		// }
		// return arr
		// }

		stats := []statsContainer{}

		// users, err := reply.IntVal("owners.[]")
		// if err != nil {
		// 	logrus.WithError(err).Warn("UpdateStats")
		// 	return
		// }

		// keys := reply.ArrayStrings("keys")
		
		// heroStats, err := r.DB.FindHeroStats(r.DB.NewSession(), heroID)
		// if err != nil {
		// 	logrus.WithError(err).Errorf("Cannot retrieve stats of hero %d", owneheroID)
		// 	return
		// }

		// statsPairs, err := r.fetchStats(&heroStats, keys)
		// if err != nil {
		// 	logrus.WithError(err)
		// 	return
		// }


		// for i := 0; i < count; i++ {
		// arr[i] = m.Get(fmt.Sprintf("%s.%d", prefix, i))
		// }
		// return arr
		// }


		stats = append(stats, statsContainer{
			OwnerID:   ownerID,
			OwnerType: 1,
			Stats:     statsPairs,
		})
	}





		numKeys, _ := event.Command.Message.IntVal(fmt.Sprintf("u.%d.s.[]", i))
		for j := 0; j < numKeys; j++ {
			pre := fmt.Sprintf("u.%d.s.%d.", i, j) //pre means prefix

			ut := reply[(pre + "ut")]
			pt := reply[(pre+ "pt")]
			key := reply[(pre + "k")]
			val := reply[(pre + "t")] // is this first value getting locked ?
			//var toUse string
			
			if val == "" {
				val = reply[(pre + "v")]
				r.changeStats(&p, key, val, ut, pt)
				logrus.Println("===UpdateStat as value 1ST=="+ key, val, ut, pt)
			}else {
				val = reply[(pre + "t")]
				r.changeStats(&p, key, val, ut, pt)	
				logrus.Println("Update as text 2nd"+ key, val, ut, pt)				
			}	
		}
		if err != nil{
			logrus.Error(err)
		}
		r.commitStats(sess, &p, heroID)

	}

	r.answer(event.Client, event.Command.PayloadID, ansUpdateStats{
		Txn: rankUpdateStats,
	})
}
