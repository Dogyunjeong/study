import * as CryptoJS from "crypto-js"

class Block {
    public index: number
    public hash: string
    public previousHash: string
    public data: string
    public timestamp: number

    constructor (
        index: number,
        hash: string,
        previousHash: string,
        data: string,
        timestamp: number,
    ) {
        this.index = index
        this.hash = hash
        this.previousHash = previousHash
        this.data = data
        this.timestamp = timestamp
    }

    static calculateBlockHash = (index: number, previousHash: string, timestamp: number, data: string): string => {
        return CryptoJS.SHA256(index + previousHash + timestamp + data).toString();
    }

    static validateStructure = (aBlock: Block) : boolean =>
        typeof aBlock.index === "number" &&
        typeof aBlock.hash === "string" &&
        typeof aBlock.previousHash === "string" &&
        typeof aBlock.timestamp === "string" &&
        typeof aBlock.data === "string" &&
}

const genesisBlock:Block = new Block(0, "20202002020002200", "", "Hello", 123456);

let blockChain: Block[] = [genesisBlock]

const getBlockChain = (): Block[] => blockChain

const getLatestBlock = (): Block => blockChain[blockChain.length -1]

const getNewTimeStamp = (): number => Math.round(new Date().getTime() /1000)

const createNewBlock = (data: string) : Block => {
    const previousBlock : Block = getLatestBlock();
    const newIndex: number = previousBlock.index + 1
    const newTimestamp: number = getNewTimeStamp();
    const newHash: string = Block.calculateBlockHash(newIndex, previousBlock.hash, newTimestamp, data)
    const newBlock: Block = new Block(newIndex, newHash, previousBlock.hash, data, newTimestamp)
    return newBlock
}

const isBlockValid = (candidateBlock: Block, previousBlock: Block): boolean => {
    if (Block.validateStructure(candidateBlock)) {
        return false
    }
    if (previousBlock.index + 1 !== candidateBlock.index) {
        return false
    }
    if (previousBlock.hash !== candidateBlock.previousHash) {
        return false
    }
    if 
}

console.log(createNewBlock("hello"), createNewBlock("bye bye"))

console.log(blockChain)