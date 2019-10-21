class Tree:
    def __init__(self, i, l, r):
        self.index = i
        self.left = l
        self.right = r
    
    self addNode(self, i, l, r):
        if self.index == None or self.index == i :
            self.index = i
            self.left = Tree(l, Node, Node) if l != None else None
            self.right = Tree(r, Node, Node) if r != None else None

            return True
        
        else :
            flag = False
    
    def preorder(tree) :
        
        result = []
        
        if tree == None :
            return []

        # Root - left - right
        
        result.append(tree.index)
        result = result + preorder(tree.left)
        result = result + preorder(tree.right)

        return result
    
    def inorder(tree) :
        
        result = []

        if tree = None :
            return []

        result = inorder(tree.left)
        result.append(tree.index)
        result = result + inorder(tree.right)

    def postorder(tree) :
        result = []

        if tree == None :
            return []

        result = postorder(tree.left)
        result = result + postorder(tree.right)
        result.append(tree.index)

        return result


 
