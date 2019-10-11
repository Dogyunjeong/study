new Vue ({
  el: '#app',
  data: {
    playerHp: 100,
    monsterHp: 100,
    log: [],
    isPlaying: false
  },
  methods: {
    startGame: function () {
      this.playerHp = 100
      this.monsterHp = 100
      this.log = []
      this.isPlaying = true
    },
    attack: function () {
      this.playerAttack(3, 10)
      this.monsterAttack()
    },
    specialAttack: function () {
      this.playerAttack(5, 15)
      this.monsterAttack()
    },
    heal: function () {
      const heal = Math.floor(Math.random() * 20)
      this.playerHp += heal
      this.log.unshift({ isPlayer: true, text: 'Player heal for ' + heal })
      this.monsterAttack
    },
    playerAttack: function (min, max) {
      const damage = this.calculateDamage(min, max)
      this.monsterHp -= damage
      this.log.unshift({
        isPlayer: true,
        text: 'Player hits Monster for' + damage
      })
      this.checkWin()
    },
    monsterAttack: function () {
      if (!this.isPlaying) {
        return
      }
      const damage = this.calculateDamage(5, 12)
      this.playerHp -= damage
      this.log.unshift({
        isPlayer: false,
        text: 'Monster hits player for' + damage
      })
      this.checkWin()
    },
    calculateDamage: function (min, max) {
      return Math.max(Math.floor(Math.random() * max) + 1, min)
    },
    checkWin: function () {
      if (this.monsterHp <= 0) {
        if (confirm('You won new game?')) {
          this.startGame()
        } else {
          this.isPlaying = false
        }
      }
      if (this.playerHp <= 0) {
        if (confirm('You Lost new game?')) {
          this.startGame()
        } else {
          this.isPlaying = false
        }
      }
    }
  }
})
