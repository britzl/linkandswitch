components {
  id: "script"
  component: "/linkandswitch/block.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "block_above"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_KINEMATIC\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"block_above\"\nmask: \"block_below\"\nmask: \"spawner\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_BOX\n    position {\n      x: 0.0\n      y: 34.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 3\n  }\n  data: 50.0\n  data: 20.0\n  data: 10.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "block_below"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_KINEMATIC\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"block_below\"\nmask: \"block_above\"\nmask: \"ground\"\nmask: \"spawner\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_BOX\n    position {\n      x: 0.0\n      y: -28.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 3\n  }\n  data: 50.0\n  data: 24.0\n  data: 10.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "down"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"down\"\nmask: \"touch\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: -100.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 2.5\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "left"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"left\"\nmask: \"touch\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: -100.0\n      y: 0.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 2.5\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "right"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"right\"\nmask: \"touch\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 100.0\n      y: 0.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 2.5\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/linkandswitch/game.atlas\"\ndefault_animation: \"blue\"\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "touch"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"touch\"\nmask: \"cursor\"\nmask: \"left\"\nmask: \"right\"\nmask: \"up\"\nmask: \"down\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 40.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "up"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"up\"\nmask: \"touch\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: 100.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 2.5\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
