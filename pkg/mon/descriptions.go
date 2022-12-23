package mon

func DescriptionPlanet(p Planet) string {
	switch p {
	case Moon:
		return "The Days and Hours of the Moon are good for embassies; voyages envoys; messages; navigation; reconciliation; love; and the acquisition of merchandise by water."
	case Mercury:
		return "The Days and Hours of Mercury are good to operate for eloquence and intelligence; promptitude in business; science and divination; wonders; apparitions; and answers regarding the future. Thou canst also operate under this Planet for thefts; writings; deceit; and merchandise."
	case Venus:
		return "The Days and Hours of Venus are good for forming friendships; for kindness and love; for joyous and pleasant undertakings, and for travelling."
	case Sun:
		return "The Days and Hours of the Sun are very good for perfecting experiments regarding temporal wealth, hope, gain, fortune, divination, the favour of princes, to dissolve hostile feeling, and to make friends."
	case Mars:
		return "In the Days and Hours of Mars thou canst make experiments regarding War; to arrive at military honour; to acquire courage; to overthrow enemies; and further to cause ruin, slaughter, cruelty, discord; to wound and to give death."
	case Jupeter:
		return "The Days and Hours of Jupiter are proper for obtaining honours, acquiring riches; contracting friendships, preserving health; and arriving at all that thou canst desire."
	case Saturn:
		return "In the Days and Hours of Saturn thou canst perform experiments to summon the Souls from Hades, but only of those who have died a natural death. Similarly on these days and hours thou canst operate to bring either good or bad fortune to buildings; to have familiar Spirits attend thee in sleep; to cause good or ill success to business, possessions, goods, seeds, fruits, and similar things, in order to acquire learning; to bring destruction and to give death, and to sow hatred and discord."
	default:
		return "-1"
	}
}
