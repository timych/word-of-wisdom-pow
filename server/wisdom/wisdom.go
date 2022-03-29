package wisdom

import (
	"math/rand"
	"time"
)

var wordsOfWisdom = []string{
	"Don't cry because it's over, smile because it happened.",
	"Be yourself; everyone else is already taken.",
	"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe",
	"The Superior Man is aware of Righteousness, the inferior man is aware of advantage.",
	"Myths which are believed in tend to become true.",
	"In all chaos there is a cosmos, in all disorder a secret order.",
	"Be courteous to all, but intimate with few, and let those few be well tried before you give them your confidence.",
	"If you break your neck, if you have nothing to eat, if your house is on fire, then you got a problem. Everything else is inconvenience.",
	"The past has no power to stop you from being present now. Only your grievance about the past can do that.",
	"There is nothing happens to any person but what was in his power to go through with.",
	"The smallest act of kindness is worth more than the grandest intention.",
	"Every problem has a gift for you in its hands.",
	"The greatest minds are capable of the greatest vices as well as of the greatest virtues.",
	"The world makes way for the man who knows where he is going.",
	"You are a product of your environment. So choose the environment that will best develop you toward your objective. Analyze your life in terms of its environment. Are the things around you helping you toward success - or are they holding you back?",
	"The cautious seldom err.",
	"The most complicated achievements of thought are possible without the assistance of consciousness.",
	"As we express our gratitude, we must never forget that the highest appreciation is not to utter words, but to live by them.",
	"The greatest way to live with honor in this world is to be what we pretend to be.",
	"Every person, all the events of your life are there because you have drawn them there. What you choose to do with them is up to you.",
	"Wisdom begins at the end.",
	"Ignorant men don't know what good they hold in their hands until they've flung it away.",
	"However many holy words you read, however many you speak, what good will they do you if you do not act on upon them?",
	"Science gives us knowledge, but only philosophy can give us wisdom.",
	"A passion for politics stems usually from an insatiable need, either for power, or for friendship and adulation, or a combination of both.",
	"Be not angry that you cannot make others as you wish them to be, since you cannot make yourself as you wish to be.",
	"No one can make you feel inferior without your consent.",
	"Appreciation can make a day, even change a life. Your willingness to put it into words is all that is necessary.",
	"Much wisdom often goes with fewest words.",
	"Joy is the best makeup.",
	"There never was a good knife made of bad steel.",
	"I love my past. I love my present. Im not ashamed of what Ive had, and Im not sad because I have it no longer.",
	"There is a difference between happiness and wisdom: he that thinks himself the happiest man is really so; but he that thinks himself the wisest is generally the greatest fool.",
	"To free us from the expectations of others, to give us back to ourselves... there lies the great, singular power of self-respect.",
	"The power of intuitive understanding will protect you from harm until the end of your days.",
	"Never explain - your friends do not need it and your enemies will not believe you anyway.",
	"Discipline is the bridge between goals and accomplishment.",
	"It is not only for what we do that we are held responsible, but also for what we do not do.",
	"All difficult things have their origin in that which is easy, and great things in that which is small.",
	"It requires wisdom to understand wisdom: the music is nothing if the audience is deaf.",
	"So is cheerfulness, or a good temper, the more it is spent, the more remains.",
	"Trust yourself. You know more than you think you do.",
	"Never promise more than you can perform.",
	"Formula for success: under promise and over deliver.",
	"The only limit to our realization of tomorrow will be our doubts of today.",
	"I think people who are creative are the luckiest people on earth. I know that there are no shortcuts, but you must keep your faith in something Greater than You, and keep doing what you love. Do what you love, and you will find the way to get it out to the world.",
	"The possibilities are numerous once we decide to act and not react.",
	"It is better to have enough ideas for some of them to be wrong, than to be always right by having no ideas at all.",
	"The most I can do for my friend is simply be his friend.",
	"He who lives in harmony with himself lives in harmony with the world.",
}

func init() {
	rand.Seed(time.Now().Unix())
}

func GetWordOfWisdom() string {
	n := rand.Int() % len(wordsOfWisdom)
	return wordsOfWisdom[n]
}
