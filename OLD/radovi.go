package OLD

import (
	"github.com/marcetin/wingcal/model"
)

func NewRadovi() map[int]model.WingVrstaRadova {

	radovi := &map[int]model.WingVrstaRadova{
		0: model.WingVrstaRadova{
			Id:             1,
			Naziv:          "PRIPREMNO ZAVRŠNI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "pripremno_zavrsni_radovi",
		},
		1: model.WingVrstaRadova{
			Id:             2,
			Naziv:          "ISTRAŽIVAČKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "istrazivacki_radovi",
		},
		2: model.WingVrstaRadova{
			Id:             3,
			Naziv:          "DEMONTAŽE I RUŠENJA",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "demontaze_i_rusenja",
		},
		3: model.WingVrstaRadova{
			Id:             4,
			Naziv:          "ZEMLJANI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "zemljani_radovi",
		},
		4: model.WingVrstaRadova{
			Id:             5,
			Naziv:          "ZIDARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "zidarski_radovi",
		},
		5: model.WingVrstaRadova{
			Id:             6,
			Naziv:          "BETONSKI I ARM. BETONSKI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "betonski_i_arm_betonski",
		},
		6: model.WingVrstaRadova{
			Id:             7,
			Naziv:          "TESARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "tesarski_radovi",
		},
		7: model.WingVrstaRadova{
			Id:             8,
			Naziv:          "POKRIVAČKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "pokrivacki_radovi",
		},
		8: model.WingVrstaRadova{
			Id:             9,
			Naziv:          "IZOLATERSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "izolaterski_radovi",
		},
		9: model.WingVrstaRadova{
			Id:             10,
			Naziv:          "GRAĐEVINSKA STOLARIJA",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "gradjevinska_stolarija",
		},
		10: model.WingVrstaRadova{
			Id:             11,
			Naziv:          "STOLARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "stolarski_radovi",
		},
		11: model.WingVrstaRadova{
			Id:             12,
			Naziv:          "BRAVARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "bravarski_radovi",
		},
		12: model.WingVrstaRadova{
			Id:             13,
			Naziv:          "LIMARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "limarski_radovi",
		},
		13: model.WingVrstaRadova{
			Id:             14,
			Naziv:          "STAKLOREZAČKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "staklorezacki_radovi",
		},
		14: model.WingVrstaRadova{
			Id:             15,
			Naziv:          "KERAMIČARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "keramicarski_radovi",
		},
		15: model.WingVrstaRadova{
			Id:             16,
			Naziv:          "TERACERSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "teracerski_radovi",
		},
		16: model.WingVrstaRadova{
			Id:             17,
			Naziv:          "KAMENOREZAČKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "kamenorezacki_radovi",
		},
		17: model.WingVrstaRadova{
			Id:             18,
			Naziv:          "PARKETARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "parketarski_radovi",
		},
		18: model.WingVrstaRadova{
			Id:             19,
			Naziv:          "PODOPOLAGAČKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "podopolagacki_radovi",
		},
		19: model.WingVrstaRadova{
			Id:             20,
			Naziv:          "TAPETARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "tapetarski_radovi",
		},
		20: model.WingVrstaRadova{
			Id:             21,
			Naziv:          "ROLETNARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "roletnarski_radovi",
		},
		21: model.WingVrstaRadova{
			Id:             22,
			Naziv:          "SUVOMONTAŽNI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "suvomontazni_radovi",
		},
		22: model.WingVrstaRadova{
			Id:             23,
			Naziv:          "GIPSARSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "gipsarski_radovi",
		},
		23: model.WingVrstaRadova{
			Id:             24,
			Naziv:          "FASADERSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "fasaderski_radovi",
		},
		24: model.WingVrstaRadova{
			Id:             25,
			Naziv:          "LIKOREZAČKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "likorezacki_radovi",
		},
		25: molerskofarbarskiradovi(),
		26: model.WingVrstaRadova{
			Id:             27,
			Naziv:          "LIVAČKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "livacki_radovi",
		},
		27: model.WingVrstaRadova{
			Id:             28,
			Naziv:          "RAZNI ZANATSKI RADOVI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "razni_zanatski_radovi",
		},
		28: model.WingVrstaRadova{
			Id:             29,
			Naziv:          "VODOVOD",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "vodovod",
		},
		29: model.WingVrstaRadova{
			Id:             30,
			Naziv:          "KANALIZACIJA",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "kanalizacija",
		},
		30: model.WingVrstaRadova{
			Id:             31,
			Naziv:          "SANITARNI UREĐAJI",
			PodvrsteRadova: pripremnozavrsniradovi(),
			Slug:           "sanitarni_uredjaji",
		},
	}
	return radovi
}
